# Table: bitbucket_repositories

This table shows data for Bitbucket Repositories.

The primary key for this table is **_cq_id**.

## Columns

| Name          | Type          |
| ------------- | ------------- |
|_cq_source_name|`utf8`|
|_cq_sync_time|`timestamp[us, tz=UTC]`|
|_cq_id (PK)|`uuid`|
|_cq_parent_id|`uuid`|
|type|`utf8`|
|full_name|`utf8`|
|links|`json`|
|name|`utf8`|
|slug|`utf8`|
|description|`utf8`|
|scm|`utf8`|
|owner|`json`|
|workspace|`json`|
|is_private|`bool`|
|project|`json`|
|fork_policy|`utf8`|
|created_on|`timestamp[us, tz=UTC]`|
|updated_on|`timestamp[us, tz=UTC]`|
|size|`int64`|
|language|`utf8`|
|has_issues|`bool`|
|has_wiki|`bool`|
|uuid|`utf8`|
|mainbranch|`json`|
|override_settings|`json`|