import { defineStore } from 'pinia';
export let startTime = (new Date() as any) / 1000 - 60 * 60 * 2;
export let endTime = (new Date() as any) / 1000;
export const useLayoutStore = defineStore('layoutOption', {
  state: () => {
    return {
      layout_option: [
        {
          x: 0, y: 0, w: 1, h: 4, i: '0',
          static: true, display: true, title: '运行时间',
          query: {
            sqls: [{ sql: '(time()-node_boot_time_seconds{instance="{macIp}"})/(60*60)' }],
            type: 'value', range: false, isValue: true, interval: 5,
            target: 'value_series', unit: 'h', float: 2
          }
        },
        {
          x: 1, y: 0, w: 1, h: 4, i: '1',
          static: true, display: true, title: 'CPU核数',
          query: {
            sqls: [{ sql: 'count(count(node_cpu_seconds_total{instance="{macIp}",mode="system"}) by (cpu))' }],
            type: 'value', range: false, isValue: true, interval: 5,
            target: 'value_series', unit: '个', float: 0
          }
        },
        {
          x: 2, y: 0, w: 2, h: 4, i: '2',
          static: true, display: true, title: '内存总量',
          query: {
            sqls: [{ sql: 'node_memory_MemTotal_bytes{instance="{macIp}"}' }],
            type: 'value', range: false, isValue: true, interval: 5,
            target: 'byte2GB_series', unit: 'GiB', float: 2,

          }
        },
        {
          x: 4, y: 0, w: 2, h: 4, i: '3',
          static: true, display: true, title: 'CPU总使用率',
          query: {
            sqls: [{ sql: '100 - (avg(irate(node_cpu_seconds_total{instance="{macIp}",mode="idle"}[5m])) * 100)' }],
            type: 'gauge', range: false, isChart: true, interval: 5,
            target: 'percent_series', unit: '%', float: 2, min: 0, max: 100,
            color: [
              [0.5, '#67e0e3'],
              [0.8, '#E6A23C'],
              [1, '#fd666d']
            ]
          }
        },
        {
          x: 6, y: 0, w: 2, h: 4, i: '4',
          static: true, display: true, title: 'CPU iowait',
          query: {
            sqls: [{ sql: 'avg(irate(node_cpu_seconds_total{instance="{macIp}",mode="iowait"}[5m])) * 100' }],
            type: 'gauge', range: false, isChart: true, interval: 5,
            target: 'percent_series', unit: '%', float: 2, min: 0, max: 100,
            color: [
              [0.3, '#67e0e3'],
              [0.5, '#E6A23C'],
              [1, '#fd666d']
            ]
          }
        },
        {
          x: 8, y: 0, w: 2, h: 4, i: '5',
          static: true, display: true, title: '内存使用率',
          query: {
            sqls: [{ sql: '(1 - (node_memory_MemAvailable_bytes{instance="{macIp}"} / (node_memory_MemTotal_bytes{instance="{macIp}"})))* 100' }],
            type: 'gauge', range: false, isChart: true, interval: 5,
            target: 'percent_series', unit: '%', float: 2, min: 0, max: 100,
            color: [
              [0.8, '#67e0e3'],
              [0.9, '#E6A23C'],
              [1, '#fd666d']
            ]
          }
        },
        {
          x: 10, y: 0, w: 3, h: 4, i: '6',
          static: true, display: true, title: '当前打开的文件描述符',
          query: {
            sqls: [{ sql: 'node_filefd_allocated{instance="{macIp}"}' }],
            type: 'gauge', range: false, isChart: true, interval: 5,
            target: 'num_series', unit: 'K', float: 2, min: 0, max: 9,
            color: [
              [0.6, '#67e0e3'],
              [0.9, '#E6A23C'],
              [1, '#fd666d']
            ]
          }
        },
        {
          x: 13, y: 0, w: 3, h: 4, i: '7',
          static: true, display: true, title: '根分区使用率',
          query: {
            sqls: [{ sql: '100 - ((node_filesystem_avail_bytes{instance="{macIp}",mountpoint="/",fstype=~"ext4|xfs"} * 100) / node_filesystem_size_bytes {instance="{macIp}",mountpoint="/",fstype=~"ext4|xfs"})' }],
            type: 'gauge', range: false, isChart: true, interval: 5,
            target: 'percent_series', unit: '%', float: 2, min: 0, max: 100,
            color: [
              [0.8, '#67e0e3'],
              [0.9, '#E6A23C'],
              [1, '#fd666d']
            ]
          }
        },
        {
          x: 0, y: 4, w: 6, h: 7, i: '8',
          static: false, display: true, title: '系统平均负载',
          query: {
            type: 'line', range: true, isChart: true, interval: 5,
            target: 'value_series', unit: '', float: 2, min: 0, max: null,
            sqls: [
              {
                sql: 'node_load1{instance="{macIp}"}',
                start: startTime,
                end: endTime,
                series_name: '1分钟'
              },
              {
                sql: 'node_load5{instance="{macIp}"}',
                start: startTime,
                end: endTime,
                series_name: '5分钟'
              },
              {
                sql: 'node_load15{instance="{macIp}"}',
                start: startTime,
                end: endTime,
                series_name: '15分钟'
              },
            ],
          }
        },
        {
          x: 6, y: 4, w: 4, h: 7, i: '9',
          static: false, display: true, title: '磁盘总空间',
          query: {
            type: 'table', range: false, isChart: false, interval: 5,
            target: 'more_query',
            sqls: [{
              sql: 'node_filesystem_size_bytes{instance="{macIp}",fstype=~"ext4|xfs"}/10^9',
              columnList: [
                { prop: "filesystem", label: '文件系统' },
                { prop: "zone", label: '分区' },
                { prop: "size", label: '空间大小(GB)' },
              ],
              columnValue: [
                { prop: 'filesystem', value: 'item.metric.fstype', type: '' },
                { prop: 'zone', value: 'item.metric.mountpoint', type: '' },
                { prop: 'size', value: 'item.value[1]', type: 'float' }
              ]
            }]
          }
        },
        {
          x: 10, y: 4, w: 6, h: 7, i: '10',
          static: false, display: true, title: '各分区可用空间',
          query: {
            type: 'table', range: false, isChart: false, interval: 5,
            target: 'more_query',
            sqls: [{
              sql: 'node_filesystem_avail_bytes {instance="{macIp}",fstype=~"ext4|xfs"}',
              columnList: [
                { prop: "filesystem", label: '文件系统' },
                { prop: "zone", label: '分区' },
                { prop: "avail", label: '可用空间(GB)' },
              ],
              columnValue: [
                { prop: 'filesystem', value: 'item.metric.fstype', type: '' },
                { prop: 'zone', value: 'item.metric.mountpoint', type: '' },
                { prop: 'avail', value: 'item.value[1]', type: 'byte', }
              ]
            },
            {
              sql: '1-(node_filesystem_free_bytes{instance="{macIp}",fstype=~"ext4|xfs"} / node_filesystem_size_bytes{instance="{macIp}",fstype=~"ext4|xfs"})',
              columnList: [
                { prop: "use", label: '使用率' },
              ],
              columnValue: [
                { prop: 'use', value: 'item.value[1]', type: 'percent', }
              ]
            }]
          }
        },
        {
          x: 0, y: 11, w: 10, h: 7, i: '11',
          static: false, display: true, title: 'cpu使用率',
          query: {
            type: 'line', range: true, isChart: true,
            target: 'percent_series', unit: '%', float: 2, min: 0, max: null,
            sqls: [
              {
                sql: 'avg(irate(node_cpu_seconds_total{instance="{macIp}",mode="system"}[1m])) by (instance)',
                start: startTime,
                end: endTime,
                step: 15,
                series_name: '系统cpu使用率',
              },
              {
                sql: 'avg(irate(node_cpu_seconds_total{instance="{macIp}",mode="user"}[1m])) by (instance)',
                start: startTime,
                end: endTime,
                step: 15,
                series_name: '用户cpu使用率',
              },
              {
                sql: 'avg(irate(node_cpu_seconds_total{instance="{macIp}",mode="idle"}[1m])) by (instance)',
                start: startTime,
                end: endTime,
                step: 15,
                series_name: '单核cpu空闲率',
              },
              {
                sql: 'avg(irate(node_cpu_seconds_total{instance="{macIp}",mode="iowait"}[1m])) by (instance)',
                start: startTime,
                end: endTime,
                step: 15,
                series_name: '磁盘io使用率',
              },
              /* {
                sql: 'irate(node_disk_io_time_seconds_total{instance="{macIp}",}[1m])',
                start: startTime,
                end: endTime,
                step: 15,
                series_name: '',
              }, */
            ]
          }
        },
        {
          x: 10, y: 11, w: 6, h: 7, i: '12',
          static: false, display: true, title: '内存信息',
          query: {
            type: 'line', range: true, isChart: true,
            target: 'byte2GB_series', unit: 'GiB', float: 2, min: 0, max: null,
            sqls: [
              {
                sql: 'node_memory_MemTotal_bytes{instance="{macIp}"}',
                start: startTime,
                end: endTime,
                step: 15,
                series_name: '总内存',
              },
              {
                sql: 'node_memory_MemTotal_bytes{instance="{macIp}"} - node_memory_MemAvailable_bytes{instance="{macIp}"}',
                start: startTime,
                end: endTime,
                step: 15,
                series_name: '已用',
              },
              {
                sql: 'node_memory_MemAvailable_bytes{instance="{macIp}"}',
                start: startTime,
                end: endTime,
                step: 15,
                series_name: '可用',
              },

            ]
          }
        },
        {
          x: 0, y: 18, w: 8, h: 7, i: '13',
          static: false, display: true, title: '磁盘读取容量',
          query: {
            type: 'line', range: true, isChart: true,
            target: 'byte2KB_series', unit: 'kB/s', float: 2, min: 0, max: null,
            sqls: [
              {
                sql: 'rate(node_disk_read_bytes_total{instance="{macIp}"}[1m])',
                start: startTime,
                end: endTime,
                step: 15,
                series_name: '读取',
              },


            ]
          }
        },
        {
          x: 8, y: 18, w: 8, h: 7, i: '14',
          static: false, display: true, title: '磁盘写入容量',
          query: {
            type: 'line', range: true, isChart: true,
            target: 'byte2KB_series', unit: 'kB/s', float: 2, min: 0, max: null,
            sqls: [
              {
                sql: 'rate(node_disk_written_bytes_total{instance="{macIp}"}[1m])',
                start: startTime,
                end: endTime,
                step: 15,
                series_name: '写入',
              },

            ]
          }
        },
        {
          x: 0, y: 25, w: 8, h: 7, i: '15',
          static: false, display: true, title: 'TCP连接情况',
          query: {
            type: 'line', range: true, isChart: true,
            target: '', unit: '', float: 2, min: 0, max: null,
            sqls: [
              {
                sql: 'node_netstat_Tcp_CurrEstab{instance="{macIp}"}',
                start: startTime,
                end: endTime,
                step: 15,
                series_name: 'ESTABLISHED',
              },
              {
                sql: 'node_sockstat_TCP_tw{instance="{macIp}"}',
                start: startTime,
                end: endTime,
                step: 15,
                series_name: 'TCP_tw',
              },
              {
                sql: 'irate(node_netstat_Tcp_ActiveOpens{instance="{macIp}"}[1m])',
                start: startTime,
                end: endTime,
                step: 15,
                series_name: 'ActiveOpens',
              },
              {
                sql: 'irate(node_netstat_Tcp_PassiveOpens{instance="{macIp}"}[1m])',
                start: startTime,
                end: endTime,
                step: 15,
                series_name: 'PassiveOpens',
              },
              {
                sql: 'node_sockstat_TCP_alloc{instance="{macIp}"}',
                start: startTime,
                end: endTime,
                step: 15,
                series_name: 'TCP_alloc',
              },
              {
                sql: 'node_sockstat_TCP_inuse{instance="{macIp}"}',
                start: startTime,
                end: endTime,
                step: 15,
                series_name: 'TCP_inuse',
              },
            ]
          }
        },
      ],
    };
  },
  /* persist: {
    enabled: true, // 开启存储
    strategies: [
      { storage: localStorage, paths: ["layout_option"] },
    ]
  }, */
  getters: {},
  actions: {
    initLayout(layout: any) {
      this.layout_option = layout;
    },
    addLayout(layout: any) {
      this.layout_option.push(layout);
    }
  }
});
