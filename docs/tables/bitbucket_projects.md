# Table: bitbucket_projects

This table shows data for Bitbucket Projects.

The primary key for this table is **uuid**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id|`uuid`|
|_cq_parent_id|`uuid`|
|type|`utf8`|
|links|`json`|
|uuid (PK)|`utf8`|
|key|`utf8`|
|owner|`json`|
|name|`utf8`|
|description|`utf8`|
|is_private|`bool`|
|created_on|`utf8`|
|updated_on|`utf8`|
|has_publicly_visible_repos|`bool`|