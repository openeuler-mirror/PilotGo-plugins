%define         debug_package %{nil}

Name:           ant-agent
Version:        3.0.0
Release:        1
Summary:        PilotGo automation plugin provides script execution and orchestration.
License:        MulanPSL-2.0
URL:            https://gitee.com/openeuler/PilotGo-plugins/automation
Source0:        ant-agent.tar.gz

BuildRequires:  systemd
Provides:       ant-agent = %{version}-%{release}

%description
PilotGo automation plugin provides script execution and orchestration.

%prep
%autosetup -p1 -n automation

%build
pushd agent
CGO_ENABLED=0 GO111MODULE=on go build -o ant-agent ./cmd/main.go
popd

%install
mkdir -p %{buildroot}/opt/PilotGo/plugin/automation/agent/log
install -D -m 0755 server/ant-agent %{buildroot}/opt/PilotGo/plugin/automation/agent
install -D -m 0644 server/ant-agent.yaml %{buildroot}/opt/PilotGo/plugin/automation/server/ant-agent.yaml
install -D -m 0644 server/scripts/ant-agent.service %{buildroot}%{_unitdir}/ant-agent.service

%post
%systemd_post ant-agent.service

%preun
%systemd_preun ant-agent.service

%postun
%systemd_postun ant-agent.service

%files
%dir /opt/PilotGo
%dir /opt/PilotGo/plugin
%dir /opt/PilotGo/plugin/automation
%dir /opt/PilotGo/plugin/automation/agent
%dir /opt/PilotGo/plugin/automation/agent/log
/opt/PilotGo/plugin/automation/agent/ant-agent
/opt/PilotGo/plugin/automation/agent/ant-agent.yaml
%{_unitdir}/ant-agent.service


%changelog
* Wed Sep 03 2025 zhanghan  <zhanghan@kylinos.cn> - 3.0.0-1
- Package init

