LOCAL_PATH := $(call my-dir)

include $(CLEAR_VARS)
LOCAL_MODULE := client-prebuilt
LOCAL_SRC_FILES := libclient.so
include $(PREBUILT_SHARED_LIBRARY)
