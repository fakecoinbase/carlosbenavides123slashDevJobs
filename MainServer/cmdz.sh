sudo docker run --net=host --rm confluentinc/cp-kafka:5.0.0 kafka-topics --create --topic new_job --partitions 4 --replication-factor 2 --if-not-exists --zookeeper localhost:32181

sudo docker run --net=host --rm confluentinc/cp-kafka:5.0.0 kafka-topics --create --topic del_job --partitions 4 --replication-factor 2 --if-not-exists --zookeeper localhost:32181


CREATE DATABASE devjobs;


drop tables companies, jobs, jobs_pivot, levels;

CREATE TABLE jobs_pivot(
    job_uuid varchar(255),
    company_uuid varchar(255),
    PRIMARY KEY(job_uuid),
    index(company_uuid)
    # FOREIGN KEY(company_uuid) REFERENCES companies(company_uuid), 
    # FOREIGN KEY(jobs_uuid) REFERENCES jobs(jobs_uuid)
);

CREATE TABLE jobs (
    job_uuid varchar(255),
    company_uuid varchar(255),
    job_title varchar(255),
    job_link varchar(255),
    job_location varchar(255),
    job_posted varchar(255),
    job_found varchar(255),
    active TINYINT(1),
    experience_level int,
    job_idx int AUTO_INCREMENT,
    PRIMARY KEY (job_uuid),
    index(job_idx),
    FOREIGN KEY(company_uuid) REFERENCES jobs_pivot(company_uuid) 
);

CREATE TABLE companies (
    company_uuid varchar(255) NOT NULL,
    company_name varchar(255) NOT NULL,
    company_cloudinary varchar(255) NOT NULL,
    PRIMARY KEY(company_uuid)
);

CREATE TABLE levels (
    id int AUTO_INCREMENT,
    job_level varchar(255),
    PRIMARY KEY (id)
);

INSERT INTO levels(job_level) VALUES ("Intern");
INSERT INTO levels(job_level) VALUES ("Entry");
INSERT INTO levels(job_level) VALUES ("Mid");
INSERT INTO levels(job_level) VALUES ("Senior");
INSERT INTO levels(job_level) VALUES ("Manager");

# INSERT INTO companies(company_uuid, name, cloudinary) 
# VALUES ("50b3dae9-0bec-456f-af6d-61a8fabe0935", "Honey", "https://res.cloudinary.com/dhxwdb3jl/image/upload/v1586121171/unnamed_wqeqel.png");

select j.job_uuid
from jobs j
left outer join jobs_pivot jp on
    jp.job_uuid = j.job_uuid
left outer join companies c on
    c.company_uuid = jp.company_uuid;

select *
from jobs inner join levels
on jobs.experience_level = levels.id
AND jobs.experience_level = 1;


select j.*, c.*, l.job_level
from companies c
inner join jobs_pivot jp on jp.company_uuid = c.company_uuid
inner join jobs j on jp.job_uuid = j.job_uuid
inner join levels l on j.experience_level = l.id;
# ORDER BY j.job_posted DESC;
# and j.job_uuid = "d7eb6e06-525b-4e93-bf3e-9e804a68748a_%_Engineering%Manager%-%iOS_%_Downtown%Los%Angeles";

select j.job_uuid, j.job_title, j.job_link, j.job_posted, j.job_found, j.job_idx, c.name, c.cloudinary, l.job_level
from companies c
inner join jobs_pivot jp on jp.company_uuid = c.company_uuid
inner join jobs j on jp.job_uuid = j.job_uuid
inner join levels l on j.experience_level = l.id
ORDER BY j.job_posted DESC;




drop tables companies, jobs, jobs_pivot, levels;