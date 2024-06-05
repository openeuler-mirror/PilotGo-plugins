package kibanaClient

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"gitee.com/openeuler/PilotGo-plugin-elk/kibanaClient/7_17_16/meta"
)

func (client *KibanaClient_v7) GetPackageInfo(ctx context.Context, pkgname string) (*meta.PackageInfo_p, error) {
	apiURL := fmt.Sprintf(meta.FleetPackageInfoAPI, pkgname)
	resp, err := client.Client.Connection.SendWithContext(ctx, http.MethodGet, apiURL, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error calling %s API: %w", meta.FleetPackageInfoAPI, err)
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	pinfo := Gjson_GetPackageInfo(bytes, "response.name", "response.policy_templates", "response.data_streams", "response.version", "response.title")
	return pinfo, nil
}
