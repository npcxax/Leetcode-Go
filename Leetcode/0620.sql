# Write your MySQL query statement below

select *
from cinema
where   id mod 2 = 1 and description!="boring"
order by rating desc