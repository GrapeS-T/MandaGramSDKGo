//go:build windows && amd64 && cgo

#include "plugin.h"
#include "_cgo_export.h"

static const PluginMetadata metadata = {
    "Plain Text Go", "1.0.0", "MandaGram Community",
    "Removes the .plain prefix from outgoing text"
};

PLUGIN_EXPORT const PluginMetadata* mandagram_plugin_metadata(void) {
    return &metadata;
}

static const char* before_send(int64_t chat_id, const char* text, uintptr_t user_data) {
    (void)chat_id;
    (void)user_data;
    return goPlainBeforeSend((char*)text);
}

MessageBeforeSendFn plain_before_send_address(void) {
    return before_send;
}

PLUGIN_EXPORT bool mandagram_plugin_init(const PluginApi* api) {
    if (!api || !api->register_message_send_hook) return false;
    return goPlainInit((void*)api) != 0;
}
