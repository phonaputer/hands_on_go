CREATE USER 'dockeruser'@'%' IDENTIFIED BY 'dockerpass';

GRANT ALL PRIVILEGES ON *.* TO 'dockeruser'@'%';
