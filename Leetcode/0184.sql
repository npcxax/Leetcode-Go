# Write your MySQL query statement below

select 
    D.name as Department, 
    E.name as Employee, 
    E.salary as Salary
from
    Employee as E,
    Department as D,
    (
        select departmentId , max(salary) as salary
        from Employee
        group by departmentId
    ) as M
where 
    E.departmentId = D.id
    and E.departmentId = M.departmentId
    and E.salary = M.salary;

select
    Department.name as Department,
    Employee.name as Employee,
    Employee.salary as salary
from 
    Employee inner join Department
on Employee.departmentId = Department.id
where (Employee.salary,Employee.departmentId)
    in (
        select max(salary),departmentId
        from Employee
        group by departmentId
    );