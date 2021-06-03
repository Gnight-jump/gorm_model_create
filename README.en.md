# gorm_model_create

#### Description
Connect to the database to automatically generate the corresponding form model

#### Software Architecture
Modify conf.yaml according to the comments
```yaml
# modify the input data as required, currently only support MySQL database
Database:
# database configuration
Host: 127.0.0.1
Port: 3306
User: the test
Password: 123321
DBName: gotest
GModel:
# store the path, the last also need a divider
StorePath: ". / dao/models/"
# do not overwrite the original file, false is not overwrite
ModelCover: true,
# The name of the stored Package
PackageName: "models"
# create table name from model
TableName:
- "screenshot"
- "omgphone"
```