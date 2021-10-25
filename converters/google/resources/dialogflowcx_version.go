// ----------------------------------------------------------------------------
//
//     ***     AUTO GENERATED CODE    ***    Type: MMv1     ***
//
// ----------------------------------------------------------------------------
//
//     This file is automatically generated by Magic Modules and manual
//     changes will be clobbered when the file is regenerated.
//
//     Please read more about how to change this file in
//     .github/CONTRIBUTING.md.
//
// ----------------------------------------------------------------------------

package google

import "reflect"

const DialogflowCXVersionAssetType string = "dialogflow.googleapis.com/Version"

func resourceConverterDialogflowCXVersion() ResourceConverter {
	return ResourceConverter{
		AssetType: DialogflowCXVersionAssetType,
		Convert:   GetDialogflowCXVersionCaiObject,
	}
}

func GetDialogflowCXVersionCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//dialogflow.googleapis.com/{{parent}}/versions/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetDialogflowCXVersionApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: DialogflowCXVersionAssetType,
			Resource: &AssetResource{
				Version:              "v3",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/dialogflow/v3/rest",
				DiscoveryName:        "Version",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetDialogflowCXVersionApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandDialogflowCXVersionDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	descriptionProp, err := expandDialogflowCXVersionDescription(d.Get("description"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("description"); !isEmptyValue(reflect.ValueOf(descriptionProp)) && (ok || !reflect.DeepEqual(v, descriptionProp)) {
		obj["description"] = descriptionProp
	}

	return obj, nil
}

func expandDialogflowCXVersionDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandDialogflowCXVersionDescription(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}