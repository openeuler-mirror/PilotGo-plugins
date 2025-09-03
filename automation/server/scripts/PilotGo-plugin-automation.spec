%define         debug_package %{nil}

Name:           PilotGo-plugin-automation
Version:        3.0.0
Release:        1
Summary:        PilotGo automation plugin provides script execution and orchestration.
License:        MulanPSL-2.0
URL:            https://gitee.com/openeuler/PilotGo-plugins/automation
Source0:        PilotGo-plugin-automation.tar.gz

BuildRequires:  systemd
Provides:       pilotgo-plugin-automation = %{version}-%{release}

%description
PilotGo automation plugin provides script execution and orchestration.

%prep
%autosetup -p1 -n automation

%build
pushd server
GO111MODULE=on go build -o PilotGo-plugin-automation ./cmd/main.go
popd

pushd web
yarn install
yarn run build
popd

%install
mkdir -p %{buildroot}/opt/PilotGo/plugin/automation/{server/log,web/dist}
install -D -m 0755 server/PilotGo-plugin-automation %{buildroot}/opt/PilotGo/plugin/automation/server
install -D -m 0644 server/automation.yml %{buildroot}/opt/PilotGo/plugin/automation/server/automation.yml
install -D -m 0644 server/scripts/PilotGo-plugin-automation.service %{buildroot}%{_unitdir}/PilotGo-plugin-automation.service
cp -rf web/dist %{buildroot}/opt/PilotGo/plugin/automation/web

%post
%systemd_post PilotGo-plugin-automation.service

%preun
%systemd_preun PilotGo-plugin-automation.service

%postun
%systemd_postun PilotGo-plugin-automation.service

%files
%dir /opt/PilotGo
%dir /opt/PilotGo/plugin
%dir /opt/PilotGo/plugin/automation
%dir /opt/PilotGo/plugin/automation/server
%dir /opt/PilotGo/plugin/automation/server/log
/opt/PilotGo/plugin/automation/server/PilotGo-plugin-automation
/opt/PilotGo/plugin/automation/server/automation.yml
%{_unitdir}/PilotGo-plugin-automation.service
/opt/PilotGo/plugin/automation/web/dist


%changelog
* Wed Sep 03 2025 zhanghan  <zhanghan@kylinos.cn> - 3.0.0-1
- Package init

