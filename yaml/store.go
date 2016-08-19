package yaml

import (
	"fmt"
	"github.com/autrilla/yaml"
	"go.mozilla.org/sops"
	"go.mozilla.org/sops/kms"
	"go.mozilla.org/sops/pgp"
	"time"
)

type YAMLStore struct {
}

func (store YAMLStore) mapSliceToTreeBranch(in yaml.MapSlice) sops.TreeBranch {
	branch := make(sops.TreeBranch, 0)
	for _, item := range in {
		branch = append(branch, sops.TreeItem{
			Key:   item.Key.(string),
			Value: store.yamlValueToTreeValue(item.Value),
		})
	}
	return branch
}

func (store YAMLStore) Load(in string) (sops.TreeBranch, error) {
	var data yaml.MapSlice
	if err := yaml.Unmarshal([]byte(in), &data); err != nil {
		return nil, fmt.Errorf("Error unmarshaling input YAML: %s", err)
	}
	for i, item := range data {
		if item.Key == "sops" {
			data = append(data[:i], data[i+1:]...)
		}
	}
	return store.mapSliceToTreeBranch(data), nil
}

func (store YAMLStore) yamlValueToTreeValue(in interface{}) interface{} {
	switch in := in.(type) {
	case map[interface{}]interface{}:
		return store.yamlMapToTreeBranch(in)
	case yaml.MapSlice:
		return store.mapSliceToTreeBranch(in)
	case []interface{}:
		return store.yamlSliceToTreeValue(in)
	default:
		return in
	}
}

func (store *YAMLStore) yamlSliceToTreeValue(in []interface{}) []interface{} {
	for i, v := range in {
		in[i] = store.yamlValueToTreeValue(v)
	}
	return in
}

func (store *YAMLStore) yamlMapToTreeBranch(in map[interface{}]interface{}) sops.TreeBranch {
	branch := make(sops.TreeBranch, 0)
	for k, v := range in {
		branch = append(branch, sops.TreeItem{
			Key:   k.(string),
			Value: store.yamlValueToTreeValue(v),
		})
	}
	return branch
}

func (store YAMLStore) treeValueToYamlValue(in interface{}) interface{} {
	switch in := in.(type) {
	case sops.TreeBranch:
		return store.treeBranchToYamlMap(in)
	default:
		return in
	}
}

func (store YAMLStore) treeBranchToYamlMap(in sops.TreeBranch) yaml.MapSlice {
	branch := make(yaml.MapSlice, 0)
	for _, item := range in {
		branch = append(branch, yaml.MapItem{
			Key:   item.Key,
			Value: store.treeValueToYamlValue(item.Value),
		})
	}
	return branch
}

func (store YAMLStore) Dump(tree sops.TreeBranch) (string, error) {
	yamlMap := store.treeBranchToYamlMap(tree)
	out, err := yaml.Marshal(yamlMap)
	if err != nil {
		return "", fmt.Errorf("Error marshaling to yaml: %s", err)
	}
	return string(out), nil
}

func (store YAMLStore) DumpWithMetadata(tree sops.TreeBranch, metadata sops.Metadata) (string, error) {
	yamlMap := store.treeBranchToYamlMap(tree)
	yamlMap = append(yamlMap, yaml.MapItem{Key: "sops", Value: metadata.ToMap()})
	out, err := yaml.Marshal(yamlMap)
	if err != nil {
		return "", fmt.Errorf("Error marshaling to yaml: %s", err)
	}
	return string(out), nil
}

func (store *YAMLStore) LoadMetadata(in string) (sops.Metadata, error) {
	var metadata sops.Metadata
	data := make(map[interface{}]interface{})
	encoded := make(map[interface{}]interface{})
	if err := yaml.Unmarshal([]byte(in), &encoded); err != nil {
		return metadata, fmt.Errorf("Error unmarshalling input yaml: %s", err)
	}

	sopsYaml, err := yaml.Marshal(encoded["sops"])
	if err != nil {
		return metadata, err
	}

	err = yaml.Unmarshal(sopsYaml, &data)
	if err != nil {
		return metadata, err
	}
	metadata.MessageAuthenticationCode = data["mac"].(string)
	lastModified, err := time.Parse(sops.DateFormat, data["lastmodified"].(string))
	if err != nil {
		return metadata, fmt.Errorf("Could not parse last modified date: %s", err)
	}
	metadata.LastModified = lastModified
	metadata.UnencryptedSuffix = data["unencrypted_suffix"].(string)
	metadata.Version = data["version"].(string)
	if k, ok := data["kms"].([]interface{}); ok {
		ks, err := store.kmsEntries(k)
		if err == nil {
			metadata.KeySources = append(metadata.KeySources, ks)
		}

	}

	if pgp, ok := data["pgp"].([]interface{}); ok {
		ks, err := store.pgpEntries(pgp)
		if err == nil {
			metadata.KeySources = append(metadata.KeySources, ks)
		}
	}
	return metadata, nil
}

func (store *YAMLStore) kmsEntries(in []interface{}) (sops.KeySource, error) {
	var keys []sops.MasterKey
	keysource := sops.KeySource{Name: "kms", Keys: keys}
	for _, v := range in {
		entry := v.(map[interface{}]interface{})
		key := &kms.KMSMasterKey{}
		key.Arn = entry["arn"].(string)
		key.EncryptedKey = entry["enc"].(string)
		role, ok := entry["role"].(string)
		if ok {
			key.Role = role
		}
		creationDate, err := time.Parse(sops.DateFormat, entry["created_at"].(string))
		if err != nil {
			return keysource, fmt.Errorf("Could not parse creation date: %s", err)
		}
		key.CreationDate = creationDate
		keysource.Keys = append(keysource.Keys, key)
	}
	return keysource, nil
}

func (store *YAMLStore) pgpEntries(in []interface{}) (sops.KeySource, error) {
	var keys []sops.MasterKey
	keysource := sops.KeySource{Name: "pgp", Keys: keys}
	for _, v := range in {
		entry := v.(map[interface{}]interface{})
		key := &pgp.GPGMasterKey{}
		key.Fingerprint = entry["fp"].(string)
		key.EncryptedKey = entry["enc"].(string)
		creationDate, err := time.Parse(sops.DateFormat, entry["created_at"].(string))
		if err != nil {
			return keysource, fmt.Errorf("Could not parse creation date: %s", err)
		}
		key.CreationDate = creationDate
		keysource.Keys = append(keysource.Keys, key)
	}
	return keysource, nil
}
