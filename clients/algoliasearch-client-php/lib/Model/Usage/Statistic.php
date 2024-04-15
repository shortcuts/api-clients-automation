<?php

// Code generated by OpenAPI Generator (https://openapi-generator.tech), manual changes will be lost - read more on https://github.com/algolia/api-clients-automation. DO NOT EDIT.

namespace Algolia\AlgoliaSearch\Model\Usage;

/**
 * Statistic Class Doc Comment.
 *
 * @category Class
 */
class Statistic
{
    /**
     * Possible values of this enum.
     */
    public const STAR = '*';

    public const SEARCH_OPERATIONS = 'search_operations';

    public const TOTAL_SEARCH_OPERATIONS = 'total_search_operations';

    public const TOTAL_SEARCH_REQUESTS = 'total_search_requests';

    public const QUERIES_OPERATIONS = 'queries_operations';

    public const MULTI_QUERIES_OPERATIONS = 'multi_queries_operations';

    public const ACL_OPERATIONS = 'acl_operations';

    public const TOTAL_ACL_OPERATIONS = 'total_acl_operations';

    public const GET_API_KEYS_OPERATIONS = 'get_api_keys_operations';

    public const GET_API_KEY_OPERATIONS = 'get_api_key_operations';

    public const ADD_API_KEY_OPERATIONS = 'add_api_key_operations';

    public const UPDATE_API_KEY_OPERATIONS = 'update_api_key_operations';

    public const DELETE_API_KEY_OPERATIONS = 'delete_api_key_operations';

    public const LIST_API_KEY_OPERATIONS = 'list_api_key_operations';

    public const INDEXING_OPERATIONS = 'indexing_operations';

    public const TOTAL_INDEXING_OPERATIONS = 'total_indexing_operations';

    public const BROWSE_OPERATIONS = 'browse_operations';

    public const CLEAR_INDEX_OPERATIONS = 'clear_index_operations';

    public const COPY_MOVE_OPERATIONS = 'copy_move_operations';

    public const DELETE_INDEX_OPERATIONS = 'delete_index_operations';

    public const GET_LOG_OPERATIONS = 'get_log_operations';

    public const GET_SETTINGS_OPERATIONS = 'get_settings_operations';

    public const SET_SETTINGS_OPERATIONS = 'set_settings_operations';

    public const LIST_INDICES_OPERATIONS = 'list_indices_operations';

    public const WAIT_TASK_OPERATIONS = 'wait_task_operations';

    public const RECORD_OPERATIONS = 'record_operations';

    public const TOTAL_RECORDS_OPERATIONS = 'total_records_operations';

    public const ADD_RECORD_OPERATIONS = 'add_record_operations';

    public const BATCH_OPERATIONS = 'batch_operations';

    public const DELETE_BY_QUERY_OPERATIONS = 'delete_by_query_operations';

    public const DELETE_RECORD_OPERATIONS = 'delete_record_operations';

    public const GET_RECORD_OPERATIONS = 'get_record_operations';

    public const PARTIAL_UPDATE_RECORD_OPERATIONS = 'partial_update_record_operations';

    public const UPDATE_RECORD_OPERATIONS = 'update_record_operations';

    public const SYNONYM_OPERATIONS = 'synonym_operations';

    public const TOTAL_SYNONYM_OPERATIONS = 'total_synonym_operations';

    public const BATCH_SYNONYM_OPERATIONS = 'batch_synonym_operations';

    public const CLEAR_SYNONYM_OPERATIONS = 'clear_synonym_operations';

    public const DELETE_SYNONYM_OPERATIONS = 'delete_synonym_operations';

    public const GET_SYNONYM_OPERATIONS = 'get_synonym_operations';

    public const QUERY_SYNONYM_OPERATIONS = 'query_synonym_operations';

    public const UPDATE_SYNONYM_OPERATIONS = 'update_synonym_operations';

    public const RULE_OPERATIONS = 'rule_operations';

    public const TOTAL_RULES_OPERATIONS = 'total_rules_operations';

    public const BATCH_RULES_OPERATIONS = 'batch_rules_operations';

    public const CLEAR_RULES_OPERATIONS = 'clear_rules_operations';

    public const DELETE_RULES_OPERATIONS = 'delete_rules_operations';

    public const GET_RULES_OPERATIONS = 'get_rules_operations';

    public const SAVE_RULES_OPERATIONS = 'save_rules_operations';

    public const SEARCH_RULES_OPERATIONS = 'search_rules_operations';

    public const TOTAL_RECOMMEND_REQUESTS = 'total_recommend_requests';

    public const TOTAL_WRITE_OPERATIONS = 'total_write_operations';

    public const TOTAL_READ_OPERATIONS = 'total_read_operations';

    public const TOTAL_OPERATIONS = 'total_operations';

    public const QUERYSUGGESTIONS_TOTAL_SEARCH_OPERATIONS = 'querysuggestions_total_search_operations';

    public const QUERYSUGGESTIONS_TOTAL_SEARCH_REQUESTS = 'querysuggestions_total_search_requests';

    public const QUERYSUGGESTIONS_TOTAL_ACL_OPERATIONS = 'querysuggestions_total_acl_operations';

    public const QUERYSUGGESTIONS_TOTAL_INDEXING_OPERATIONS = 'querysuggestions_total_indexing_operations';

    public const QUERYSUGGESTIONS_TOTAL_RECORDS_OPERATIONS = 'querysuggestions_total_records_operations';

    public const QUERYSUGGESTIONS_TOTAL_SYNONYM_OPERATIONS = 'querysuggestions_total_synonym_operations';

    public const QUERYSUGGESTIONS_TOTAL_RULES_OPERATIONS = 'querysuggestions_total_rules_operations';

    public const QUERYSUGGESTIONS_TOTAL_WRITE_OPERATIONS = 'querysuggestions_total_write_operations';

    public const QUERYSUGGESTIONS_TOTAL_READ_OPERATIONS = 'querysuggestions_total_read_operations';

    public const QUERYSUGGESTIONS_TOTAL_OPERATIONS = 'querysuggestions_total_operations';

    public const AVG_PROCESSING_TIME = 'avg_processing_time';

    public const _90P_PROCESSING_TIME = '90p_processing_time';

    public const _99P_PROCESSING_TIME = '99p_processing_time';

    public const QUERIES_ABOVE_LAST_MS_PROCESSING_TIME = 'queries_above_last_ms_processing_time';

    public const RECORDS = 'records';

    public const DATA_SIZE = 'data_size';

    public const FILE_SIZE = 'file_size';

    public const MAX_QPS = 'max_qps';

    public const REGION_MAX_QPS = 'region_max_qps';

    public const TOTAL_MAX_QPS = 'total_max_qps';

    public const USED_SEARCH_CAPACITY = 'used_search_capacity';

    public const AVG_USED_SEARCH_CAPACITY = 'avg_used_search_capacity';

    public const REGION_USED_SEARCH_CAPACITY = 'region_used_search_capacity';

    public const REGION_AVG_USED_SEARCH_CAPACITY = 'region_avg_used_search_capacity';

    public const TOTAL_USED_SEARCH_CAPACITY = 'total_used_search_capacity';

    public const TOTAL_AVG_USED_SEARCH_CAPACITY = 'total_avg_used_search_capacity';

    public const DEGRADED_QUERIES_SSD_USED_QUERIES_IMPACTED = 'degraded_queries_ssd_used_queries_impacted';

    public const DEGRADED_QUERIES_SSD_USED_SECONDS_IMPACTED = 'degraded_queries_ssd_used_seconds_impacted';

    public const DEGRADED_QUERIES_MAX_CAPACITY_QUERIES_IMPACTED = 'degraded_queries_max_capacity_queries_impacted';

    public const DEGRADED_QUERIES_MAX_CAPACITY_SECONDS_IMPACTED = 'degraded_queries_max_capacity_seconds_impacted';

    /**
     * Gets allowable values of the enum.
     *
     * @return string[]
     */
    public static function getAllowableEnumValues()
    {
        return [
            self::STAR,
            self::SEARCH_OPERATIONS,
            self::TOTAL_SEARCH_OPERATIONS,
            self::TOTAL_SEARCH_REQUESTS,
            self::QUERIES_OPERATIONS,
            self::MULTI_QUERIES_OPERATIONS,
            self::ACL_OPERATIONS,
            self::TOTAL_ACL_OPERATIONS,
            self::GET_API_KEYS_OPERATIONS,
            self::GET_API_KEY_OPERATIONS,
            self::ADD_API_KEY_OPERATIONS,
            self::UPDATE_API_KEY_OPERATIONS,
            self::DELETE_API_KEY_OPERATIONS,
            self::LIST_API_KEY_OPERATIONS,
            self::INDEXING_OPERATIONS,
            self::TOTAL_INDEXING_OPERATIONS,
            self::BROWSE_OPERATIONS,
            self::CLEAR_INDEX_OPERATIONS,
            self::COPY_MOVE_OPERATIONS,
            self::DELETE_INDEX_OPERATIONS,
            self::GET_LOG_OPERATIONS,
            self::GET_SETTINGS_OPERATIONS,
            self::SET_SETTINGS_OPERATIONS,
            self::LIST_INDICES_OPERATIONS,
            self::WAIT_TASK_OPERATIONS,
            self::RECORD_OPERATIONS,
            self::TOTAL_RECORDS_OPERATIONS,
            self::ADD_RECORD_OPERATIONS,
            self::BATCH_OPERATIONS,
            self::DELETE_BY_QUERY_OPERATIONS,
            self::DELETE_RECORD_OPERATIONS,
            self::GET_RECORD_OPERATIONS,
            self::PARTIAL_UPDATE_RECORD_OPERATIONS,
            self::UPDATE_RECORD_OPERATIONS,
            self::SYNONYM_OPERATIONS,
            self::TOTAL_SYNONYM_OPERATIONS,
            self::BATCH_SYNONYM_OPERATIONS,
            self::CLEAR_SYNONYM_OPERATIONS,
            self::DELETE_SYNONYM_OPERATIONS,
            self::GET_SYNONYM_OPERATIONS,
            self::QUERY_SYNONYM_OPERATIONS,
            self::UPDATE_SYNONYM_OPERATIONS,
            self::RULE_OPERATIONS,
            self::TOTAL_RULES_OPERATIONS,
            self::BATCH_RULES_OPERATIONS,
            self::CLEAR_RULES_OPERATIONS,
            self::DELETE_RULES_OPERATIONS,
            self::GET_RULES_OPERATIONS,
            self::SAVE_RULES_OPERATIONS,
            self::SEARCH_RULES_OPERATIONS,
            self::TOTAL_RECOMMEND_REQUESTS,
            self::TOTAL_WRITE_OPERATIONS,
            self::TOTAL_READ_OPERATIONS,
            self::TOTAL_OPERATIONS,
            self::QUERYSUGGESTIONS_TOTAL_SEARCH_OPERATIONS,
            self::QUERYSUGGESTIONS_TOTAL_SEARCH_REQUESTS,
            self::QUERYSUGGESTIONS_TOTAL_ACL_OPERATIONS,
            self::QUERYSUGGESTIONS_TOTAL_INDEXING_OPERATIONS,
            self::QUERYSUGGESTIONS_TOTAL_RECORDS_OPERATIONS,
            self::QUERYSUGGESTIONS_TOTAL_SYNONYM_OPERATIONS,
            self::QUERYSUGGESTIONS_TOTAL_RULES_OPERATIONS,
            self::QUERYSUGGESTIONS_TOTAL_WRITE_OPERATIONS,
            self::QUERYSUGGESTIONS_TOTAL_READ_OPERATIONS,
            self::QUERYSUGGESTIONS_TOTAL_OPERATIONS,
            self::AVG_PROCESSING_TIME,
            self::_90P_PROCESSING_TIME,
            self::_99P_PROCESSING_TIME,
            self::QUERIES_ABOVE_LAST_MS_PROCESSING_TIME,
            self::RECORDS,
            self::DATA_SIZE,
            self::FILE_SIZE,
            self::MAX_QPS,
            self::REGION_MAX_QPS,
            self::TOTAL_MAX_QPS,
            self::USED_SEARCH_CAPACITY,
            self::AVG_USED_SEARCH_CAPACITY,
            self::REGION_USED_SEARCH_CAPACITY,
            self::REGION_AVG_USED_SEARCH_CAPACITY,
            self::TOTAL_USED_SEARCH_CAPACITY,
            self::TOTAL_AVG_USED_SEARCH_CAPACITY,
            self::DEGRADED_QUERIES_SSD_USED_QUERIES_IMPACTED,
            self::DEGRADED_QUERIES_SSD_USED_SECONDS_IMPACTED,
            self::DEGRADED_QUERIES_MAX_CAPACITY_QUERIES_IMPACTED,
            self::DEGRADED_QUERIES_MAX_CAPACITY_SECONDS_IMPACTED,
        ];
    }
}