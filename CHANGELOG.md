# Changelog

All notable changes to this project will be documented in this file.

The format is based on [Keep a Changelog](https://keepachangelog.com/en/1.0.0/),
[markdownlint](https://dlaa.me/markdownlint/),
and this project adheres to [Semantic Versioning](https://semver.org/spec/v2.0.0.html).

# [0.6.2] - 2022-01-19

### Added to 0.6.2

- Fixing variable mis reference and added better create design pattern

# [0.6.1] - 2022-01-19

### Added to 0.6.1

- Fixing defaults for role grants to be false for impersonator and grant privileges

# [0.6.0] - 2021-12-05

### Added to 0.6.0

- Added better database connect control
  - Close session after executing sql to reduce likely hood of getting too many sessions error
  - Removed the connection requirement to specify a database for the client, so that it generalizes terraform's connection
  - Added database connection variables to control the provider wit
    - max_open_conns
    - max_idle_conns
- Updated base view resource
  - Added better query to fetch information on the base view created
  - Changed the id of the object to the catalog id in denodo
  - Added custom VQL input ability
- Updated derived view resource
  - Added better query to fetch information on the base view created
  - Removed looping on a directory feature
  - Added custom VQL input ability

# [0.5.9] - 2021-08-14

### Added to 0.5.9

- Reverting to previous create pattern in 0.5.7.
- Correcting metadata fetching when object is created.

# [0.5.8] - 2021-08-14

### Added to 0.5.8

- Reverting to previous create pattern.

# [0.5.7] - 2021-08-14

### Added to 0.5.7

- Correcting signature to 8 inputs.

## [0.5.6] - 2021-08-14

### Added to 0.5.6

- Correcting data source database to generate the SQL correctly.

## [0.5.5] - 2021-08-13

### Added to 0.5.5

- Fixed documentation typo

## [0.5.4] - 2021-08-13

### Added to 0.5.4

- Fixed required bug and set to optional

## [0.5.3] - 2021-08-13

### Added to 0.5.3

- Fixed base view defaulting to Null if parameter is not passed

## [0.5.2] - 2021-06-13

### Added to 0.5.2

- Fixing documentation typo

## [0.5.1] - 2021-06-13

### Added to 0.5.1

- Updated documentation

## [0.5.0] - 2021-06-13

### Added to 0.5.0

- Added data source to get metadata from any object in metadata catalog.
- Added `denodo_dervived_view` resource to create dervived views from files.

## [0.4.2] - 2021-06-09

### Added to 0.4.2

- Fixed some provisioning bugs for setting incorrect values in statefile
- Fixed destroy bugs for
    - denodo_jdbc_data_source
    - denodo_base_view
    - denodo_database_folder
    - denodo_database_role

## [0.4.1] - 2021-06-07

### Added to 0.4.1

- Fixed some provisioning bugs
- Updated documentation

## [0.4.0] - 2021-06-06

### Added to 0.4.0

- Updated database resource to add additional configuration on creeate of database.

## [0.3.0] - 2021-06-06

### Added to 0.3.0

- Updated role to be database_role

## [0.2.1] - 2021-06-06

### Added to 0.2.1

- Updated documentation
- Data source user_for_query_optimization corrected to use_for_query_optimization

## [0.2.0] - 2021-06-06

### Added to 0.2.0

- Updated documentation

## [0.1.0] - 2021-06-06

### Added to 0.1.0

- Updated documentation and variables


## [0.0.0] - 2021-06-06

### Added to 0.0.0

- Initial release
