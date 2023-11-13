# Write your MySQL query statement below

select 
    s1.id - 1 as id,
    s1.student
from
    seat as s1
where 
    s1.id MOD 2 = 0 
union
select 
    s2.id + 1 as id,
    s2.student
from 
    seat as s2
where 
    s2.id MOD 2 = 1
    and s2.id != (
        select 
            max(s3.id)
        from 
            seat as s3
    )
union
select
    s4.id as id,
    s4.student
from 
    seat as s4
where
    s4.id MOD 2 = 1
    and s4.id = (
        select
            max(s5.id)
        from 
            seat as s5
    )
order by id