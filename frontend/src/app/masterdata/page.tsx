"use client";

import { useMasterData } from "@/app/hooks/masterData/useMasterData";

const MasterDataExample: React.FC = () => {
  const { departments, majors, loading, error, reload } = useMasterData();

  return (
    <div>
      <h1>Master Data</h1>
      {loading && <p>Loading...</p>}
      {error && <p style={{ color: "red" }}>{error}</p>}

      <h2>Departments</h2>
      <ul>
        {departments.map((department) => (
          <li key={department.department_id}>
            {department.name} (ID: {department.department_id})
          </li>
        ))}
      </ul>

      <h2>Majors</h2>
      <ul>
        {majors.map((major) => (
          <li key={major.major_id}>
            {major.name} (ID: {major.major_id}, Department ID:{" "}
            {major.department_id})
          </li>
        ))}
      </ul>

      <button onClick={reload}>Reload</button>
    </div>
  );
};

export default MasterDataExample;
