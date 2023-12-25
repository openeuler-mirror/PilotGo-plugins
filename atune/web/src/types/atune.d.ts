export interface Task {
  id: number;
  task_name: string;
  command: string;
  task_status: string;
  create_time: string;
  update_time: string;
  results: [{}];
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
