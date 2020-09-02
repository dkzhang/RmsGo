package resourceSM

/*
月度、季度、年度计量

先给出区间范围[a,b]

（1）筛选出时间交集的项目
(projectStatus==Archived && CreateAt > b && a > UpdatedAt) ||
(projectStatus!=Archived && CreateAt > b)

(2)将项目的分配记录（时点记录），转换为区间记录（或节点连续记录，更难）

(3)筛选出与给定区间有交集的记录，调整为给定区间内记录

(4)汇总计算该项目在给定区间内的计量结果

(5)按项目长、部门、All，分别进行汇总

*/
