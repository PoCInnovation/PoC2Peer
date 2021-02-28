package com.example.android.uamp.media.library;

import java.lang.System;

@kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 2, d1 = {"\u0000\u0010\n\u0000\n\u0002\u0010\b\n\u0002\b\u0004\n\u0002\u0010\u000e\n\u0000\"\u000e\u0010\u0000\u001a\u00020\u0001X\u0086T\u00a2\u0006\u0002\n\u0000\"\u000e\u0010\u0002\u001a\u00020\u0001X\u0086T\u00a2\u0006\u0002\n\u0000\"\u000e\u0010\u0003\u001a\u00020\u0001X\u0086T\u00a2\u0006\u0002\n\u0000\"\u000e\u0010\u0004\u001a\u00020\u0001X\u0086T\u00a2\u0006\u0002\n\u0000\"\u000e\u0010\u0005\u001a\u00020\u0006X\u0082T\u00a2\u0006\u0002\n\u0000\u00a8\u0006\u0007"}, d2 = {"STATE_CREATED", "", "STATE_ERROR", "STATE_INITIALIZED", "STATE_INITIALIZING", "TAG", "", "common_debug"})
public final class MusicSourceKt {
    
    /**
     * State indicating the source was created, but no initialization has performed.
     */
    public static final int STATE_CREATED = 1;
    
    /**
     * State indicating initialization of the source is in progress.
     */
    public static final int STATE_INITIALIZING = 2;
    
    /**
     * State indicating the source has been initialized and is ready to be used.
     */
    public static final int STATE_INITIALIZED = 3;
    
    /**
     * State indicating an error has occurred.
     */
    public static final int STATE_ERROR = 4;
    private static final java.lang.String TAG = "MusicSource";
}