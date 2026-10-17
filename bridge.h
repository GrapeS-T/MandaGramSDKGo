#ifndef MANDAGRAM_GO_BRIDGE_H
#define MANDAGRAM_GO_BRIDGE_H

#include "include/mandagram_plugin.h"

_Static_assert(sizeof(void*) == 8, "MandaGram Go SDK requires a 64-bit target");
_Static_assert(sizeof(PluginApi) == 53 * sizeof(void*), "Unexpected PluginApi layout");
_Static_assert(sizeof(HookParam) == 40, "Unexpected HookParam layout");

static inline int32_t mg_register_hook(const PluginApi* api, const char* class_name, const char* method_name, const char* method_descriptor, HookBeforeFn before_cb, HookAfterFn after_cb) {
    return api->register_hook(class_name, method_name, method_descriptor, before_cb, after_cb);
}

static inline void mg_log_info(const PluginApi* api, const char* tag, const char* msg) {
    api->log_info(tag, msg);
}

static inline void mg_log_error(const PluginApi* api, const char* tag, const char* msg) {
    api->log_error(tag, msg);
}

static inline int32_t mg_get_storage_dir(const PluginApi* api, const char* plugin_name, char* out_buf, uint32_t max_len) {
    return api->get_storage_dir(plugin_name, out_buf, max_len);
}

static inline int32_t mg_get_cache_dir(const PluginApi* api, const char* plugin_name, char* out_buf, uint32_t max_len) {
    return api->get_cache_dir(plugin_name, out_buf, max_len);
}

static inline int32_t mg_resolve_asset(const PluginApi* api, const char* plugin_name, const char* asset_name, char* out_buf, uint32_t max_len) {
    return api->resolve_asset(plugin_name, asset_name, out_buf, max_len);
}

static inline int64_t mg_get_current_user_id(const PluginApi* api) {
    return api->get_current_user_id();
}

static inline int64_t mg_get_active_chat_id(const PluginApi* api) {
    return api->get_active_chat_id();
}

static inline int32_t mg_execute_tdlib_query(const PluginApi* api, const char* json_query, char* out_buf, uint32_t max_len) {
    return api->execute_tdlib_query(json_query, out_buf, max_len);
}

static inline jobject mg_create_composable_lambda(const PluginApi* api, ComposableFn cb, uintptr_t user_data) {
    return api->create_composable_lambda(cb, user_data);
}

static inline jobject mg_create_scoped_lambda(const PluginApi* api, ScopedComposableFn cb, uintptr_t user_data) {
    return api->create_scoped_lambda(cb, user_data);
}

static inline void mg_replace_composable(const PluginApi* api, HookParam* param, ComposableFn cb, uintptr_t user_data) {
    api->replace_composable(param, cb, user_data);
}

static inline void mg_composer_start_group(const PluginApi* api, jobject composer, int32_t key) {
    api->composer_start_group(composer, key);
}

static inline void mg_composer_end_group(const PluginApi* api, jobject composer) {
    api->composer_end_group(composer);
}

static inline jobject mg_composer_get_remembered(const PluginApi* api, jobject composer) {
    return api->composer_get_remembered(composer);
}

static inline void mg_composer_set_remembered(const PluginApi* api, jobject composer, jobject val) {
    api->composer_set_remembered(composer, val);
}

static inline bool mg_composer_is_skipping(const PluginApi* api, jobject composer) {
    return api->composer_is_skipping(composer);
}

static inline int32_t mg_register_update_hook(const PluginApi* api, UpdateHookFn cb, uintptr_t user_data) {
    return api->register_update_hook(cb, user_data);
}

static inline int32_t mg_register_message_send_hook(const PluginApi* api, MessageBeforeSendFn before_cb, MessageAfterSendFn after_cb, uintptr_t user_data) {
    return api->register_message_send_hook(before_cb, after_cb, user_data);
}

static inline int32_t mg_register_message_receive_hook(const PluginApi* api, MessageReceiveFn cb, uintptr_t user_data) {
    return api->register_message_receive_hook(cb, user_data);
}

static inline int32_t mg_register_chat_action(const PluginApi* api, const char* id, const char* label, ChatActionClickFn click_cb, uintptr_t user_data) {
    return api->register_chat_action(id, label, click_cb, user_data);
}

static inline jobject mg_get_object_field(const PluginApi* api, jobject obj, const char* field_name) {
    return api->get_object_field(obj, field_name);
}

static inline void mg_set_object_field(const PluginApi* api, jobject obj, const char* field_name, jobject val) {
    api->set_object_field(obj, field_name, val);
}

static inline int32_t mg_get_int_field(const PluginApi* api, jobject obj, const char* field_name) {
    return api->get_int_field(obj, field_name);
}

static inline void mg_set_int_field(const PluginApi* api, jobject obj, const char* field_name, int32_t val) {
    api->set_int_field(obj, field_name, val);
}

static inline int64_t mg_get_long_field(const PluginApi* api, jobject obj, const char* field_name) {
    return api->get_long_field(obj, field_name);
}

static inline void mg_set_long_field(const PluginApi* api, jobject obj, const char* field_name, int64_t val) {
    api->set_long_field(obj, field_name, val);
}

static inline bool mg_get_bool_field(const PluginApi* api, jobject obj, const char* field_name) {
    return api->get_bool_field(obj, field_name);
}

static inline void mg_set_bool_field(const PluginApi* api, jobject obj, const char* field_name, bool val) {
    api->set_bool_field(obj, field_name, val);
}

static inline int32_t mg_get_string_utf(const PluginApi* api, jobject jstr, char* out_buf, uint32_t max_len) {
    return api->get_string_utf(jstr, out_buf, max_len);
}

static inline jobject mg_create_string_utf(const PluginApi* api, const char* str) {
    return api->create_string_utf(str);
}

static inline jobject mg_get_attached_field(const PluginApi* api, jobject obj, const char* key) {
    return api->get_attached_field(obj, key);
}

static inline void mg_set_attached_field(const PluginApi* api, jobject obj, const char* key, jobject val) {
    api->set_attached_field(obj, key, val);
}

static inline jobject mg_get_state_value(const PluginApi* api, jobject state_obj) {
    return api->get_state_value(state_obj);
}

static inline void mg_set_state_value(const PluginApi* api, jobject state_obj, jobject val) {
    api->set_state_value(state_obj, val);
}

static inline jobject mg_modifier_empty(const PluginApi* api) {
    return api->modifier_empty();
}

static inline jobject mg_modifier_size(const PluginApi* api, jobject mod, float width_dp, float height_dp) {
    return api->modifier_size(mod, width_dp, height_dp);
}

static inline jobject mg_modifier_width(const PluginApi* api, jobject mod, float width_dp) {
    return api->modifier_width(mod, width_dp);
}

static inline jobject mg_modifier_height(const PluginApi* api, jobject mod, float height_dp) {
    return api->modifier_height(mod, height_dp);
}

static inline jobject mg_modifier_fill_max_width(const PluginApi* api, jobject mod) {
    return api->modifier_fill_max_width(mod);
}

static inline jobject mg_modifier_fill_max_height(const PluginApi* api, jobject mod) {
    return api->modifier_fill_max_height(mod);
}

static inline jobject mg_modifier_padding(const PluginApi* api, jobject mod, float start_dp, float top_dp, float end_dp, float bottom_dp) {
    return api->modifier_padding(mod, start_dp, top_dp, end_dp, bottom_dp);
}

static inline jobject mg_modifier_background(const PluginApi* api, jobject mod, uint32_t argb, float corner_radius_dp) {
    return api->modifier_background(mod, argb, corner_radius_dp);
}

static inline jobject mg_modifier_border(const PluginApi* api, jobject mod, float width_dp, uint32_t argb, float corner_radius_dp) {
    return api->modifier_border(mod, width_dp, argb, corner_radius_dp);
}

static inline jobject mg_modifier_clickable(const PluginApi* api, jobject mod, jobject on_click_lambda) {
    return api->modifier_clickable(mod, on_click_lambda);
}

static inline jobject mg_modifier_weight(const PluginApi* api, jobject scope, jobject mod, float weight) {
    return api->modifier_weight(scope, mod, weight);
}

static inline void mg_compose_row(const PluginApi* api, jobject composer, jobject mod, jobject content_lambda) {
    api->compose_row(composer, mod, content_lambda);
}

static inline void mg_compose_column(const PluginApi* api, jobject composer, jobject mod, jobject content_lambda) {
    api->compose_column(composer, mod, content_lambda);
}

static inline void mg_compose_box(const PluginApi* api, jobject composer, jobject mod, jobject content_lambda) {
    api->compose_box(composer, mod, content_lambda);
}

static inline void mg_compose_spacer(const PluginApi* api, jobject composer, jobject mod) {
    api->compose_spacer(composer, mod);
}

static inline void mg_compose_text(const PluginApi* api, jobject composer, const char* text, jobject mod, uint32_t argb, float font_size_sp, bool is_bold, int32_t max_lines) {
    api->compose_text(composer, text, mod, argb, font_size_sp, is_bold, max_lines);
}

static inline void mg_compose_image(const PluginApi* api, jobject composer, jobject bitmap, jobject mod) {
    api->compose_image(composer, bitmap, mod);
}

static inline jobject mg_get_avatar_bitmap(const PluginApi* api, const char* path) {
    return api->get_avatar_bitmap(path);
}

#endif
