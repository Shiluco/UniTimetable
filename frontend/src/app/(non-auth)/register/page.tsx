"use client";

import { LoginLayout } from "@/app/components/layout/loginLayout";
import { RegisterForm } from "@/app/components/registerForm/registerForm";
import { fetchDepartments, fetchMajors } from "@/service/masterDataService";
import { Department, Major } from "@/types/masterData";
import { useEffect, useState } from "react";

export default function LoginPage()
{
  const [departments, setDepartments] = useState<Department[]>([]);
  const [majors, setMajors] = useState<Major[]>([]);
  useEffect(() => {
    const fetchData = async () => {
      const departmentsData = await fetchDepartments();
      setDepartments(departmentsData);
      const majorsData = await fetchMajors();
      setMajors(majorsData);
    };
    fetchData();
  }, []);

  return (
    <LoginLayout>
      <RegisterForm
        departmentOptions={departments.map((department) => ({ label: department.name, value: department.department_id.toString() }))}
        majorOptions={majors.map((major) => ({ label: major.name, value: major.major_id.toString() }))}
        yearOptions={[
          { label: "1年", value: "1" },
          { label: "2年", value: "2" },
          { label: "3年", value: "3" },
          { label: "4年", value: "4" },
        ]}
      />
    </LoginLayout>
  );
}
