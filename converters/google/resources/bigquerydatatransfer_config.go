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

import (
	"context"
	"fmt"
	"reflect"

	"github.com/hashicorp/terraform-plugin-sdk/v2/helper/schema"
)

var sensitiveParams = []string{"secret_access_key"}

func sensitiveParamCustomizeDiff(_ context.Context, diff *schema.ResourceDiff, v interface{}) error {
	for _, sp := range sensitiveParams {
		mapLabel := diff.Get("params." + sp).(string)
		authLabel := diff.Get("sensitive_params.0." + sp).(string)
		if mapLabel != "" && authLabel != "" {
			return fmt.Errorf("Sensitive param [%s] cannot be set in both `params` and the `sensitive_params` block.", sp)
		}
	}
	return nil
}

const BigqueryDataTransferConfigAssetType string = "bigquerydatatransfer.googleapis.com/Config"

func resourceConverterBigqueryDataTransferConfig() ResourceConverter {
	return ResourceConverter{
		AssetType: BigqueryDataTransferConfigAssetType,
		Convert:   GetBigqueryDataTransferConfigCaiObject,
	}
}

func GetBigqueryDataTransferConfigCaiObject(d TerraformResourceData, config *Config) ([]Asset, error) {
	name, err := assetName(d, config, "//bigquerydatatransfer.googleapis.com/{{name}}")
	if err != nil {
		return []Asset{}, err
	}
	if obj, err := GetBigqueryDataTransferConfigApiObject(d, config); err == nil {
		return []Asset{{
			Name: name,
			Type: BigqueryDataTransferConfigAssetType,
			Resource: &AssetResource{
				Version:              "v1",
				DiscoveryDocumentURI: "https://www.googleapis.com/discovery/v1/apis/bigquerydatatransfer/v1/rest",
				DiscoveryName:        "Config",
				Data:                 obj,
			},
		}}, nil
	} else {
		return []Asset{}, err
	}
}

func GetBigqueryDataTransferConfigApiObject(d TerraformResourceData, config *Config) (map[string]interface{}, error) {
	obj := make(map[string]interface{})
	displayNameProp, err := expandBigqueryDataTransferConfigDisplayName(d.Get("display_name"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("display_name"); !isEmptyValue(reflect.ValueOf(displayNameProp)) && (ok || !reflect.DeepEqual(v, displayNameProp)) {
		obj["displayName"] = displayNameProp
	}
	destinationDatasetIdProp, err := expandBigqueryDataTransferConfigDestinationDatasetId(d.Get("destination_dataset_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("destination_dataset_id"); !isEmptyValue(reflect.ValueOf(destinationDatasetIdProp)) && (ok || !reflect.DeepEqual(v, destinationDatasetIdProp)) {
		obj["destinationDatasetId"] = destinationDatasetIdProp
	}
	dataSourceIdProp, err := expandBigqueryDataTransferConfigDataSourceId(d.Get("data_source_id"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_source_id"); !isEmptyValue(reflect.ValueOf(dataSourceIdProp)) && (ok || !reflect.DeepEqual(v, dataSourceIdProp)) {
		obj["dataSourceId"] = dataSourceIdProp
	}
	scheduleProp, err := expandBigqueryDataTransferConfigSchedule(d.Get("schedule"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("schedule"); !isEmptyValue(reflect.ValueOf(scheduleProp)) && (ok || !reflect.DeepEqual(v, scheduleProp)) {
		obj["schedule"] = scheduleProp
	}
	scheduleOptionsProp, err := expandBigqueryDataTransferConfigScheduleOptions(d.Get("schedule_options"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("schedule_options"); !isEmptyValue(reflect.ValueOf(scheduleOptionsProp)) && (ok || !reflect.DeepEqual(v, scheduleOptionsProp)) {
		obj["scheduleOptions"] = scheduleOptionsProp
	}
	emailPreferencesProp, err := expandBigqueryDataTransferConfigEmailPreferences(d.Get("email_preferences"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("email_preferences"); !isEmptyValue(reflect.ValueOf(emailPreferencesProp)) && (ok || !reflect.DeepEqual(v, emailPreferencesProp)) {
		obj["emailPreferences"] = emailPreferencesProp
	}
	notificationPubsubTopicProp, err := expandBigqueryDataTransferConfigNotificationPubsubTopic(d.Get("notification_pubsub_topic"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("notification_pubsub_topic"); !isEmptyValue(reflect.ValueOf(notificationPubsubTopicProp)) && (ok || !reflect.DeepEqual(v, notificationPubsubTopicProp)) {
		obj["notificationPubsubTopic"] = notificationPubsubTopicProp
	}
	dataRefreshWindowDaysProp, err := expandBigqueryDataTransferConfigDataRefreshWindowDays(d.Get("data_refresh_window_days"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("data_refresh_window_days"); !isEmptyValue(reflect.ValueOf(dataRefreshWindowDaysProp)) && (ok || !reflect.DeepEqual(v, dataRefreshWindowDaysProp)) {
		obj["dataRefreshWindowDays"] = dataRefreshWindowDaysProp
	}
	disabledProp, err := expandBigqueryDataTransferConfigDisabled(d.Get("disabled"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("disabled"); !isEmptyValue(reflect.ValueOf(disabledProp)) && (ok || !reflect.DeepEqual(v, disabledProp)) {
		obj["disabled"] = disabledProp
	}
	paramsProp, err := expandBigqueryDataTransferConfigParams(d.Get("params"), d, config)
	if err != nil {
		return nil, err
	} else if v, ok := d.GetOkExists("params"); !isEmptyValue(reflect.ValueOf(paramsProp)) && (ok || !reflect.DeepEqual(v, paramsProp)) {
		obj["params"] = paramsProp
	}

	return resourceBigqueryDataTransferConfigEncoder(d, config, obj)
}

func resourceBigqueryDataTransferConfigEncoder(d TerraformResourceData, meta interface{}, obj map[string]interface{}) (map[string]interface{}, error) {
	paramMap, ok := obj["params"]
	if !ok {
		paramMap = make(map[string]string)
	}

	var params map[string]string
	params = paramMap.(map[string]string)

	for _, sp := range sensitiveParams {
		if auth, _ := d.GetOkExists("sensitive_params.0." + sp); auth != "" {
			params[sp] = auth.(string)
		}
	}

	obj["params"] = params

	return obj, nil
}

func expandBigqueryDataTransferConfigDisplayName(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDestinationDatasetId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDataSourceId(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigSchedule(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigScheduleOptions(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedDisableAutoScheduling, err := expandBigqueryDataTransferConfigScheduleOptionsDisableAutoScheduling(original["disable_auto_scheduling"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedDisableAutoScheduling); val.IsValid() && !isEmptyValue(val) {
		transformed["disableAutoScheduling"] = transformedDisableAutoScheduling
	}

	transformedStartTime, err := expandBigqueryDataTransferConfigScheduleOptionsStartTime(original["start_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedStartTime); val.IsValid() && !isEmptyValue(val) {
		transformed["startTime"] = transformedStartTime
	}

	transformedEndTime, err := expandBigqueryDataTransferConfigScheduleOptionsEndTime(original["end_time"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEndTime); val.IsValid() && !isEmptyValue(val) {
		transformed["endTime"] = transformedEndTime
	}

	return transformed, nil
}

func expandBigqueryDataTransferConfigScheduleOptionsDisableAutoScheduling(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigScheduleOptionsStartTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigScheduleOptionsEndTime(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigEmailPreferences(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	l := v.([]interface{})
	if len(l) == 0 || l[0] == nil {
		return nil, nil
	}
	raw := l[0]
	original := raw.(map[string]interface{})
	transformed := make(map[string]interface{})

	transformedEnableFailureEmail, err := expandBigqueryDataTransferConfigEmailPreferencesEnableFailureEmail(original["enable_failure_email"], d, config)
	if err != nil {
		return nil, err
	} else if val := reflect.ValueOf(transformedEnableFailureEmail); val.IsValid() && !isEmptyValue(val) {
		transformed["enableFailureEmail"] = transformedEnableFailureEmail
	}

	return transformed, nil
}

func expandBigqueryDataTransferConfigEmailPreferencesEnableFailureEmail(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigNotificationPubsubTopic(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDataRefreshWindowDays(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigDisabled(v interface{}, d TerraformResourceData, config *Config) (interface{}, error) {
	return v, nil
}

func expandBigqueryDataTransferConfigParams(v interface{}, d TerraformResourceData, config *Config) (map[string]string, error) {
	if v == nil {
		return map[string]string{}, nil
	}
	m := make(map[string]string)
	for k, val := range v.(map[string]interface{}) {
		m[k] = val.(string)
	}
	return m, nil
}
