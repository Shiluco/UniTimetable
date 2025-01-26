// 学部型
export type Department = {
  department_id: number; // 学部ID
  name: string; // 学部名
};

// 学科型
export type Major = {
  major_id: number; // 学科ID
  department_id: number; // 所属学部ID
  name: string; // 学科名
};
