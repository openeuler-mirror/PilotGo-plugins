export interface Task {
  id: number;
  task_name: string;
  command: string;
  task_status: string;
  create_time: string;
  update_time: string;
  results: Array;
  tune_id: int;
  tune: Atune;
}

export interface Atune {
  id: number;
  description: string;
  create_time: string;
  update_time: string;
  tuneName: string;
  custom_name: string;
  workDir: string;
  prepare: string;
  tune: string;
  restore: string;
  note: string;
}

// 画布矩形的参数配置
export interface RectConfig {
  x: number;
  y: number;
  width: number;
  height: number;
  fill: string;
  stroke: string;
  shadowBlur: number;
  cornerRadius: number;
}

// 导出接口字面量
export type TaskArray = Task[];
export type AtuneArray = Atune[];

// *接口api返回结果约束不含data
export interface Result {
  code: number;
  msg: string;
}

// *接口api返回结果含有page信息
export interface ReaultData<T = any> extends Result {
  data: Task[] | Atune[];
  ok?: string;
  page: number;
  size: number;
  total: number;
}
