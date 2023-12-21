export interface Task {
  id: number;
  name: string;
  atune: string;
  state: number;
  createTime: string;
  updateTime: string;
}

export interface Atune {
  id: number;
  tuneName: string;
  workDir: string;
  prepare: string;
  tune: string;
  restore: string;
  note: string;
}
