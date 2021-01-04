package com.example.android.uamp.media;

import java.lang.System;

@kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u00002\n\u0002\u0018\u0002\n\u0002\u0010\u0000\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0004\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0010\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0010\t\n\u0002\b\u0003\u0018\u0000 \u00122\u00020\u0001:\u0001\u0012B\u000f\b\u0002\u0012\u0006\u0010\u0002\u001a\u00020\u0003\u00a2\u0006\u0002\u0010\u0004J\b\u0010\t\u001a\u0004\u0018\u00010\nJ!\u0010\u000b\u001a\u00020\f2\u0006\u0010\r\u001a\u00020\u000e2\u0006\u0010\u000f\u001a\u00020\u0010H\u0086@\u00f8\u0001\u0000\u00a2\u0006\u0002\u0010\u0011R\u0011\u0010\u0002\u001a\u00020\u0003\u00a2\u0006\b\n\u0000\u001a\u0004\b\u0005\u0010\u0006R\u000e\u0010\u0007\u001a\u00020\bX\u0082\u000e\u00a2\u0006\u0002\n\u0000\u0082\u0002\u0004\n\u0002\b\u0019\u00a8\u0006\u0013"}, d2 = {"Lcom/example/android/uamp/media/PersistentStorage;", "", "context", "Landroid/content/Context;", "(Landroid/content/Context;)V", "getContext", "()Landroid/content/Context;", "preferences", "Landroid/content/SharedPreferences;", "loadRecentSong", "Landroid/support/v4/media/MediaBrowserCompat$MediaItem;", "saveRecentSong", "", "description", "Landroid/support/v4/media/MediaDescriptionCompat;", "position", "", "(Landroid/support/v4/media/MediaDescriptionCompat;JLkotlin/coroutines/Continuation;)Ljava/lang/Object;", "Companion", "common_debug"})
public final class PersistentStorage {
    
    /**
     * Store any data which must persist between restarts, such as the most recently played song.
     */
    private android.content.SharedPreferences preferences;
    @org.jetbrains.annotations.NotNull()
    private final android.content.Context context = null;
    private static volatile com.example.android.uamp.media.PersistentStorage instance;
    public static final com.example.android.uamp.media.PersistentStorage.Companion Companion = null;
    
    @org.jetbrains.annotations.Nullable()
    public final java.lang.Object saveRecentSong(@org.jetbrains.annotations.NotNull()
    android.support.v4.media.MediaDescriptionCompat description, long position, @org.jetbrains.annotations.NotNull()
    kotlin.coroutines.Continuation<? super kotlin.Unit> p2) {
        return null;
    }
    
    @org.jetbrains.annotations.Nullable()
    public final android.support.v4.media.MediaBrowserCompat.MediaItem loadRecentSong() {
        return null;
    }
    
    @org.jetbrains.annotations.NotNull()
    public final android.content.Context getContext() {
        return null;
    }
    
    private PersistentStorage(android.content.Context context) {
        super();
    }
    
    @kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000\u001a\n\u0002\u0018\u0002\n\u0002\u0010\u0000\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0000\b\u0086\u0003\u0018\u00002\u00020\u0001B\u0007\b\u0002\u00a2\u0006\u0002\u0010\u0002J\u000e\u0010\u0005\u001a\u00020\u00042\u0006\u0010\u0006\u001a\u00020\u0007R\u0010\u0010\u0003\u001a\u0004\u0018\u00010\u0004X\u0082\u000e\u00a2\u0006\u0002\n\u0000\u00a8\u0006\b"}, d2 = {"Lcom/example/android/uamp/media/PersistentStorage$Companion;", "", "()V", "instance", "Lcom/example/android/uamp/media/PersistentStorage;", "getInstance", "context", "Landroid/content/Context;", "common_debug"})
    public static final class Companion {
        
        @org.jetbrains.annotations.NotNull()
        public final com.example.android.uamp.media.PersistentStorage getInstance(@org.jetbrains.annotations.NotNull()
        android.content.Context context) {
            return null;
        }
        
        private Companion() {
            super();
        }
    }
}