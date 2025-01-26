-- departments
CREATE TABLE IF NOT EXISTS `departments` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- majors
CREATE TABLE IF NOT EXISTS `majors` (
  `id` int NOT NULL AUTO_INCREMENT,
  `name` varchar(255) NOT NULL,
  `description` text,
  `department_id` int NOT NULL,
  PRIMARY KEY (`id`),
  UNIQUE KEY `name` (`name`),
  FOREIGN KEY (`department_id`) REFERENCES `departments` (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4;

-- テストデータの挿入
INSERT INTO departments (name, description) VALUES
('工学部', '工学に関する学部です'),
('理学部', '理学に関する学部です');

INSERT INTO majors (name, description, department_id) VALUES
('情報工学科', '情報工学に関する学科です', 1),
('電気工学科', '電気工学に関する学科です', 1),
('物理学科', '物理学に関する学科です', 2),
('化学科', '化学に関する学科です', 2); 