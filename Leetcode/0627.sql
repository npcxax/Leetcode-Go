# Write your MySQL query statement below

update salary
set sex = char(ascii(sex)^ascii('m')^ascii('f'))