package kibanaClient

import (
	"context"
	"fmt"
	"io"
	"net/http"

	"gitee.com/openeuler/PilotGo-plugin-elk/kibanaClient/meta"
)

func (client *KibanaClient) GetPackageInfo(ctx context.Context, pkgname, pkgversion string) (*meta.PackageInfo_p, error) {
	apiURL := fmt.Sprintf(meta.FleetPackageInfoAPI, pkgname, pkgversion)
	resp, err := client.Client.Connection.SendWithContext(ctx, http.MethodGet, apiURL, nil, nil, nil)
	if err != nil {
		return nil, fmt.Errorf("error calling %s API: %w", meta.FleetPackageInfoAPI, err)
	}
	defer resp.Body.Close()
	bytes, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}

	pinfo := Gjson_GetPackageInfo(bytes, "item.name", "item.policy_templates", "item.data_streams")
	return pinfo, nil
}
