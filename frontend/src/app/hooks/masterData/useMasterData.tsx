import { useState, useEffect } from "react";
import { fetchDepartments, fetchMajors } from "@/service/masterDataService";
import { Department, Major } from "@/types/masterData";

export const useMasterData = () => {
  const [departments, setDepartments] = useState<Department[]>([]);
  const [majors, setMajors] = useState<Major[]>([]);
  const [loading, setLoading] = useState<boolean>(false);
  const [error, setError] = useState<string | null>(null);

  const fetchMasterData = async () => {
    try {
      setLoading(true);
      setError(null);

      // 学部データを取得
      const fetchedDepartments = await fetchDepartments();
      setDepartments(fetchedDepartments);

      // 学部ごとに学科データを取得
      const allMajors: Major[] = [];
      for (const department of fetchedDepartments) {
        const fetchedMajors = await fetchMajors({
          department_id: department.department_id,
        });
        allMajors.push(...fetchedMajors); // 全学科をマージ
      }
      setMajors(allMajors);
    } catch (err) {
      console.error("Error loading master data:", err);
      setError("データの取得に失敗しました。再度お試しください。");
    } finally {
      setLoading(false);
    }
  };

  useEffect(() => {
    fetchMasterData();
  }, []);

  return {
    departments,
    majors,
    loading,
    error,
    reload: fetchMasterData, // 再読み込み用の関数をエクスポート
  };
};
