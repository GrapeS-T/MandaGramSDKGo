//go:build windows && amd64 && cgo

// Package mandagram provides borrowed-pointer bindings to the 53-slot MandaGram PluginApi.
// The API table and every retained callback must remain valid for their entire native lifetime.
package mandagram

/*
#cgo CFLAGS: -std=c11
#include "bridge.h"
*/
import "C"

import "unsafe"

// Object is an opaque host-owned jobject. It is not a Go object or a stable identity.
type Object unsafe.Pointer

// HookParam mirrors the canonical native structure. Only use it over native memory;
// its pointer fields must never contain Go pointers. Padding is part of the ABI.
type HookParam struct {
	ThisPtr unsafe.Pointer
	Args *unsafe.Pointer
	ArgsCount uint32
	Padding uint32
	Result unsafe.Pointer
	ReturnEarly uint8
	HasResult uint8
	Reserved uint16
	Padding2 uint32
}

// HookBeforeFn is a native C function address with the signature in mandagram_plugin.h.
// Obtain it through a C trampoline, never from a Go function value.
type HookBeforeFn unsafe.Pointer

// HookAfterFn is a native C function address with the signature in mandagram_plugin.h.
// Obtain it through a C trampoline, never from a Go function value.
type HookAfterFn unsafe.Pointer

// ComposableFn is a native C function address with the signature in mandagram_plugin.h.
// Obtain it through a C trampoline, never from a Go function value.
type ComposableFn unsafe.Pointer

// ScopedComposableFn is a native C function address with the signature in mandagram_plugin.h.
// Obtain it through a C trampoline, never from a Go function value.
type ScopedComposableFn unsafe.Pointer

// UpdateHookFn is a native C function address with the signature in mandagram_plugin.h.
// Obtain it through a C trampoline, never from a Go function value.
type UpdateHookFn unsafe.Pointer

// MessageBeforeSendFn is a native C function address with the signature in mandagram_plugin.h.
// Obtain it through a C trampoline, never from a Go function value.
type MessageBeforeSendFn unsafe.Pointer

// MessageAfterSendFn is a native C function address with the signature in mandagram_plugin.h.
// Obtain it through a C trampoline, never from a Go function value.
type MessageAfterSendFn unsafe.Pointer

// MessageReceiveFn is a native C function address with the signature in mandagram_plugin.h.
// Obtain it through a C trampoline, never from a Go function value.
type MessageReceiveFn unsafe.Pointer

// ChatActionClickFn is a native C function address with the signature in mandagram_plugin.h.
// Obtain it through a C trampoline, never from a Go function value.
type ChatActionClickFn unsafe.Pointer

// API borrows the host function table; it does not own or copy it.
// Calls require a non-nil table and the corresponding non-nil slot.
// Strings are NUL-terminated C memory; output buffers are caller-owned.
type API struct {
	ptr *C.PluginApi
}

// FromPointer borrows a canonical PluginApi pointer supplied to plugin initialization.
// The caller must ensure it belongs to the matching engine revision.
func FromPointer(ptr unsafe.Pointer) API {
	return API{ptr: (*C.PluginApi)(ptr)}
}

// RegisterHook invokes PluginApi.register_hook without taking ownership of its arguments.
func (api API) RegisterHook(className unsafe.Pointer, methodName unsafe.Pointer, methodDescriptor unsafe.Pointer, beforeCb HookBeforeFn, afterCb HookAfterFn) int32 {
	return int32(C.mg_register_hook(api.ptr, (*C.char)(className), (*C.char)(methodName), (*C.char)(methodDescriptor), C.HookBeforeFn(beforeCb), C.HookAfterFn(afterCb)))
}

// LogInfo invokes PluginApi.log_info without taking ownership of its arguments.
func (api API) LogInfo(tag unsafe.Pointer, msg unsafe.Pointer) {
	C.mg_log_info(api.ptr, (*C.char)(tag), (*C.char)(msg))
}

// LogError invokes PluginApi.log_error without taking ownership of its arguments.
func (api API) LogError(tag unsafe.Pointer, msg unsafe.Pointer) {
	C.mg_log_error(api.ptr, (*C.char)(tag), (*C.char)(msg))
}

// GetStorageDir invokes PluginApi.get_storage_dir without taking ownership of its arguments.
func (api API) GetStorageDir(pluginName unsafe.Pointer, outBuf unsafe.Pointer, maxLen uint32) int32 {
	return int32(C.mg_get_storage_dir(api.ptr, (*C.char)(pluginName), (*C.char)(outBuf), C.uint32_t(maxLen)))
}

// GetCacheDir invokes PluginApi.get_cache_dir without taking ownership of its arguments.
func (api API) GetCacheDir(pluginName unsafe.Pointer, outBuf unsafe.Pointer, maxLen uint32) int32 {
	return int32(C.mg_get_cache_dir(api.ptr, (*C.char)(pluginName), (*C.char)(outBuf), C.uint32_t(maxLen)))
}

// ResolveAsset invokes PluginApi.resolve_asset without taking ownership of its arguments.
func (api API) ResolveAsset(pluginName unsafe.Pointer, assetName unsafe.Pointer, outBuf unsafe.Pointer, maxLen uint32) int32 {
	return int32(C.mg_resolve_asset(api.ptr, (*C.char)(pluginName), (*C.char)(assetName), (*C.char)(outBuf), C.uint32_t(maxLen)))
}

// GetCurrentUserID invokes PluginApi.get_current_user_id without taking ownership of its arguments.
func (api API) GetCurrentUserID() int64 {
	return int64(C.mg_get_current_user_id(api.ptr))
}

// GetActiveChatID invokes PluginApi.get_active_chat_id without taking ownership of its arguments.
func (api API) GetActiveChatID() int64 {
	return int64(C.mg_get_active_chat_id(api.ptr))
}

// ExecuteTDLibQuery invokes PluginApi.execute_tdlib_query without taking ownership of its arguments.
func (api API) ExecuteTDLibQuery(jSONQuery unsafe.Pointer, outBuf unsafe.Pointer, maxLen uint32) int32 {
	return int32(C.mg_execute_tdlib_query(api.ptr, (*C.char)(jSONQuery), (*C.char)(outBuf), C.uint32_t(maxLen)))
}

// CreateComposableLambda invokes PluginApi.create_composable_lambda without taking ownership of its arguments.
func (api API) CreateComposableLambda(cb ComposableFn, userData uintptr) Object {
	return Object(C.mg_create_composable_lambda(api.ptr, C.ComposableFn(cb), C.uintptr_t(userData)))
}

// CreateScopedLambda invokes PluginApi.create_scoped_lambda without taking ownership of its arguments.
func (api API) CreateScopedLambda(cb ScopedComposableFn, userData uintptr) Object {
	return Object(C.mg_create_scoped_lambda(api.ptr, C.ScopedComposableFn(cb), C.uintptr_t(userData)))
}

// ReplaceComposable invokes PluginApi.replace_composable without taking ownership of its arguments.
func (api API) ReplaceComposable(param *HookParam, cb ComposableFn, userData uintptr) {
	C.mg_replace_composable(api.ptr, (*C.HookParam)(unsafe.Pointer(param)), C.ComposableFn(cb), C.uintptr_t(userData))
}

// ComposerStartGroup invokes PluginApi.composer_start_group without taking ownership of its arguments.
func (api API) ComposerStartGroup(composer Object, key int32) {
	C.mg_composer_start_group(api.ptr, C.jobject(composer), C.int32_t(key))
}

// ComposerEndGroup invokes PluginApi.composer_end_group without taking ownership of its arguments.
func (api API) ComposerEndGroup(composer Object) {
	C.mg_composer_end_group(api.ptr, C.jobject(composer))
}

// ComposerGetRemembered invokes PluginApi.composer_get_remembered without taking ownership of its arguments.
func (api API) ComposerGetRemembered(composer Object) Object {
	return Object(C.mg_composer_get_remembered(api.ptr, C.jobject(composer)))
}

// ComposerSetRemembered invokes PluginApi.composer_set_remembered without taking ownership of its arguments.
func (api API) ComposerSetRemembered(composer Object, val Object) {
	C.mg_composer_set_remembered(api.ptr, C.jobject(composer), C.jobject(val))
}

// ComposerIsSkipping invokes PluginApi.composer_is_skipping without taking ownership of its arguments.
func (api API) ComposerIsSkipping(composer Object) bool {
	return bool(C.mg_composer_is_skipping(api.ptr, C.jobject(composer)))
}

// RegisterUpdateHook invokes PluginApi.register_update_hook without taking ownership of its arguments.
func (api API) RegisterUpdateHook(cb UpdateHookFn, userData uintptr) int32 {
	return int32(C.mg_register_update_hook(api.ptr, C.UpdateHookFn(cb), C.uintptr_t(userData)))
}

// RegisterMessageSendHook invokes PluginApi.register_message_send_hook without taking ownership of its arguments.
func (api API) RegisterMessageSendHook(beforeCb MessageBeforeSendFn, afterCb MessageAfterSendFn, userData uintptr) int32 {
	return int32(C.mg_register_message_send_hook(api.ptr, C.MessageBeforeSendFn(beforeCb), C.MessageAfterSendFn(afterCb), C.uintptr_t(userData)))
}

// RegisterMessageReceiveHook invokes PluginApi.register_message_receive_hook without taking ownership of its arguments.
func (api API) RegisterMessageReceiveHook(cb MessageReceiveFn, userData uintptr) int32 {
	return int32(C.mg_register_message_receive_hook(api.ptr, C.MessageReceiveFn(cb), C.uintptr_t(userData)))
}

// RegisterChatAction invokes PluginApi.register_chat_action without taking ownership of its arguments.
func (api API) RegisterChatAction(iD unsafe.Pointer, label unsafe.Pointer, clickCb ChatActionClickFn, userData uintptr) int32 {
	return int32(C.mg_register_chat_action(api.ptr, (*C.char)(iD), (*C.char)(label), C.ChatActionClickFn(clickCb), C.uintptr_t(userData)))
}

// GetObjectField invokes PluginApi.get_object_field without taking ownership of its arguments.
func (api API) GetObjectField(obj Object, fieldName unsafe.Pointer) Object {
	return Object(C.mg_get_object_field(api.ptr, C.jobject(obj), (*C.char)(fieldName)))
}

// SetObjectField invokes PluginApi.set_object_field without taking ownership of its arguments.
func (api API) SetObjectField(obj Object, fieldName unsafe.Pointer, val Object) {
	C.mg_set_object_field(api.ptr, C.jobject(obj), (*C.char)(fieldName), C.jobject(val))
}

// GetIntField invokes PluginApi.get_int_field without taking ownership of its arguments.
func (api API) GetIntField(obj Object, fieldName unsafe.Pointer) int32 {
	return int32(C.mg_get_int_field(api.ptr, C.jobject(obj), (*C.char)(fieldName)))
}

// SetIntField invokes PluginApi.set_int_field without taking ownership of its arguments.
func (api API) SetIntField(obj Object, fieldName unsafe.Pointer, val int32) {
	C.mg_set_int_field(api.ptr, C.jobject(obj), (*C.char)(fieldName), C.int32_t(val))
}

// GetLongField invokes PluginApi.get_long_field without taking ownership of its arguments.
func (api API) GetLongField(obj Object, fieldName unsafe.Pointer) int64 {
	return int64(C.mg_get_long_field(api.ptr, C.jobject(obj), (*C.char)(fieldName)))
}

// SetLongField invokes PluginApi.set_long_field without taking ownership of its arguments.
func (api API) SetLongField(obj Object, fieldName unsafe.Pointer, val int64) {
	C.mg_set_long_field(api.ptr, C.jobject(obj), (*C.char)(fieldName), C.int64_t(val))
}

// GetBoolField invokes PluginApi.get_bool_field without taking ownership of its arguments.
func (api API) GetBoolField(obj Object, fieldName unsafe.Pointer) bool {
	return bool(C.mg_get_bool_field(api.ptr, C.jobject(obj), (*C.char)(fieldName)))
}

// SetBoolField invokes PluginApi.set_bool_field without taking ownership of its arguments.
func (api API) SetBoolField(obj Object, fieldName unsafe.Pointer, val bool) {
	C.mg_set_bool_field(api.ptr, C.jobject(obj), (*C.char)(fieldName), C.bool(val))
}

// GetStringUTF invokes PluginApi.get_string_utf without taking ownership of its arguments.
func (api API) GetStringUTF(jstr Object, outBuf unsafe.Pointer, maxLen uint32) int32 {
	return int32(C.mg_get_string_utf(api.ptr, C.jobject(jstr), (*C.char)(outBuf), C.uint32_t(maxLen)))
}

// CreateStringUTF invokes PluginApi.create_string_utf without taking ownership of its arguments.
func (api API) CreateStringUTF(str unsafe.Pointer) Object {
	return Object(C.mg_create_string_utf(api.ptr, (*C.char)(str)))
}

// GetAttachedField invokes PluginApi.get_attached_field without taking ownership of its arguments.
func (api API) GetAttachedField(obj Object, key unsafe.Pointer) Object {
	return Object(C.mg_get_attached_field(api.ptr, C.jobject(obj), (*C.char)(key)))
}

// SetAttachedField invokes PluginApi.set_attached_field without taking ownership of its arguments.
func (api API) SetAttachedField(obj Object, key unsafe.Pointer, val Object) {
	C.mg_set_attached_field(api.ptr, C.jobject(obj), (*C.char)(key), C.jobject(val))
}

// GetStateValue invokes PluginApi.get_state_value without taking ownership of its arguments.
func (api API) GetStateValue(stateObj Object) Object {
	return Object(C.mg_get_state_value(api.ptr, C.jobject(stateObj)))
}

// SetStateValue invokes PluginApi.set_state_value without taking ownership of its arguments.
func (api API) SetStateValue(stateObj Object, val Object) {
	C.mg_set_state_value(api.ptr, C.jobject(stateObj), C.jobject(val))
}

// ModifierEmpty invokes PluginApi.modifier_empty without taking ownership of its arguments.
func (api API) ModifierEmpty() Object {
	return Object(C.mg_modifier_empty(api.ptr))
}

// ModifierSize invokes PluginApi.modifier_size without taking ownership of its arguments.
func (api API) ModifierSize(mod Object, widthDp float32, heightDp float32) Object {
	return Object(C.mg_modifier_size(api.ptr, C.jobject(mod), C.float(widthDp), C.float(heightDp)))
}

// ModifierWidth invokes PluginApi.modifier_width without taking ownership of its arguments.
func (api API) ModifierWidth(mod Object, widthDp float32) Object {
	return Object(C.mg_modifier_width(api.ptr, C.jobject(mod), C.float(widthDp)))
}

// ModifierHeight invokes PluginApi.modifier_height without taking ownership of its arguments.
func (api API) ModifierHeight(mod Object, heightDp float32) Object {
	return Object(C.mg_modifier_height(api.ptr, C.jobject(mod), C.float(heightDp)))
}

// ModifierFillMaxWidth invokes PluginApi.modifier_fill_max_width without taking ownership of its arguments.
func (api API) ModifierFillMaxWidth(mod Object) Object {
	return Object(C.mg_modifier_fill_max_width(api.ptr, C.jobject(mod)))
}

// ModifierFillMaxHeight invokes PluginApi.modifier_fill_max_height without taking ownership of its arguments.
func (api API) ModifierFillMaxHeight(mod Object) Object {
	return Object(C.mg_modifier_fill_max_height(api.ptr, C.jobject(mod)))
}

// ModifierPadding invokes PluginApi.modifier_padding without taking ownership of its arguments.
func (api API) ModifierPadding(mod Object, startDp float32, topDp float32, endDp float32, bottomDp float32) Object {
	return Object(C.mg_modifier_padding(api.ptr, C.jobject(mod), C.float(startDp), C.float(topDp), C.float(endDp), C.float(bottomDp)))
}

// ModifierBackground invokes PluginApi.modifier_background without taking ownership of its arguments.
func (api API) ModifierBackground(mod Object, argb uint32, cornerRadiusDp float32) Object {
	return Object(C.mg_modifier_background(api.ptr, C.jobject(mod), C.uint32_t(argb), C.float(cornerRadiusDp)))
}

// ModifierBorder invokes PluginApi.modifier_border without taking ownership of its arguments.
func (api API) ModifierBorder(mod Object, widthDp float32, argb uint32, cornerRadiusDp float32) Object {
	return Object(C.mg_modifier_border(api.ptr, C.jobject(mod), C.float(widthDp), C.uint32_t(argb), C.float(cornerRadiusDp)))
}

// ModifierClickable invokes PluginApi.modifier_clickable without taking ownership of its arguments.
func (api API) ModifierClickable(mod Object, onClickLambda Object) Object {
	return Object(C.mg_modifier_clickable(api.ptr, C.jobject(mod), C.jobject(onClickLambda)))
}

// ModifierWeight invokes PluginApi.modifier_weight without taking ownership of its arguments.
func (api API) ModifierWeight(scope Object, mod Object, weight float32) Object {
	return Object(C.mg_modifier_weight(api.ptr, C.jobject(scope), C.jobject(mod), C.float(weight)))
}

// ComposeRow invokes PluginApi.compose_row without taking ownership of its arguments.
func (api API) ComposeRow(composer Object, mod Object, contentLambda Object) {
	C.mg_compose_row(api.ptr, C.jobject(composer), C.jobject(mod), C.jobject(contentLambda))
}

// ComposeColumn invokes PluginApi.compose_column without taking ownership of its arguments.
func (api API) ComposeColumn(composer Object, mod Object, contentLambda Object) {
	C.mg_compose_column(api.ptr, C.jobject(composer), C.jobject(mod), C.jobject(contentLambda))
}

// ComposeBox invokes PluginApi.compose_box without taking ownership of its arguments.
func (api API) ComposeBox(composer Object, mod Object, contentLambda Object) {
	C.mg_compose_box(api.ptr, C.jobject(composer), C.jobject(mod), C.jobject(contentLambda))
}

// ComposeSpacer invokes PluginApi.compose_spacer without taking ownership of its arguments.
func (api API) ComposeSpacer(composer Object, mod Object) {
	C.mg_compose_spacer(api.ptr, C.jobject(composer), C.jobject(mod))
}

// ComposeText invokes PluginApi.compose_text without taking ownership of its arguments.
func (api API) ComposeText(composer Object, text unsafe.Pointer, mod Object, argb uint32, fontSizeSp float32, isBold bool, maxLines int32) {
	C.mg_compose_text(api.ptr, C.jobject(composer), (*C.char)(text), C.jobject(mod), C.uint32_t(argb), C.float(fontSizeSp), C.bool(isBold), C.int32_t(maxLines))
}

// ComposeImage invokes PluginApi.compose_image without taking ownership of its arguments.
func (api API) ComposeImage(composer Object, bitmap Object, mod Object) {
	C.mg_compose_image(api.ptr, C.jobject(composer), C.jobject(bitmap), C.jobject(mod))
}

// GetAvatarBitmap invokes PluginApi.get_avatar_bitmap without taking ownership of its arguments.
func (api API) GetAvatarBitmap(path unsafe.Pointer) Object {
	return Object(C.mg_get_avatar_bitmap(api.ptr, (*C.char)(path)))
}

