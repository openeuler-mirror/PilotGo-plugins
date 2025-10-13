%define         debug_package %{nil}

Name:           ismp-automation
Version:        3.0.0
Release:        1
Summary:        PilotGo automation plugin provides script execution and orchestration.
License:        MulanPSL-2.0
URL:            https://gitee.com/openeuler/PilotGo-plugins/automation
Source0:        ismp-automation.tar.gz

BuildRequires:  systemd
Provides:       ismp-automation = %{version}-%{release}

%description
PilotGo automation plugin provides script execution and orchestration.

%prep
%autosetup -p1 -n automation

%build
pushd server
CGO_ENABLED=0 GO111MODULE=on go build -o ismp-automation ./cmd/main.go
popd

pushd web
yarn install
yarn run build
popd

%install
mkdir -p %{buildroot}/opt/PilotGo/plugin/automation/{server/log,web/dist}
install -D -m 0755 server/ismp-automation %{buildroot}/opt/PilotGo/plugin/automation/server
install -D -m 0644 server/automation.yaml %{buildroot}/opt/PilotGo/plugin/automation/server/automation.yaml
install -D -m 0644 server/scripts/ismp-automation.service %{buildroot}%{_unitdir}/ismp-automation.service
cp -rf web/dist %{buildroot}/opt/PilotGo/plugin/automation/web

%post
%systemd_post ismp-automation.service

%preun
%systemd_preun ismp-automation.service

%postun
%systemd_postun ismp-automation.service

%files
%dir /opt/PilotGo
%dir /opt/PilotGo/plugin
%dir /opt/PilotGo/plugin/automation
%dir /opt/PilotGo/plugin/automation/server
%dir /opt/PilotGo/plugin/automation/server/log
/opt/PilotGo/plugin/automation/server/ismp-automation
/opt/PilotGo/plugin/automation/server/automation.yaml
%{_unitdir}/ismp-automation.service
/opt/PilotGo/plugin/automation/web/dist


%changelog
* Wed Sep 03 2025 zhanghan  <zhanghan@kylinos.cn> - 3.0.0-1
- Package init

