package com.example.android.uamp.utils;

import java.lang.System;

/**
 * Static methods used to inject classes needed for various Activities and Fragments.
 */
@kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u00000\n\u0002\u0018\u0002\n\u0002\u0010\u0000\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0010\u000e\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\b\u00c6\u0002\u0018\u00002\u00020\u0001B\u0007\b\u0002\u00a2\u0006\u0002\u0010\u0002J\u000e\u0010\u0003\u001a\u00020\u00042\u0006\u0010\u0005\u001a\u00020\u0006J\u0016\u0010\u0007\u001a\u00020\b2\u0006\u0010\u0005\u001a\u00020\u00062\u0006\u0010\t\u001a\u00020\nJ\u0010\u0010\u000b\u001a\u00020\f2\u0006\u0010\u0005\u001a\u00020\u0006H\u0002J\u000e\u0010\r\u001a\u00020\u000e2\u0006\u0010\u0005\u001a\u00020\u0006\u00a8\u0006\u000f"}, d2 = {"Lcom/example/android/uamp/utils/InjectorUtils;", "", "()V", "provideMainActivityViewModel", "Lcom/example/android/uamp/viewmodels/MainActivityViewModel$Factory;", "context", "Landroid/content/Context;", "provideMediaItemFragmentViewModel", "Lcom/example/android/uamp/viewmodels/MediaItemFragmentViewModel$Factory;", "mediaId", "", "provideMusicServiceConnection", "Lcom/example/android/uamp/common/MusicServiceConnection;", "provideNowPlayingFragmentViewModel", "Lcom/example/android/uamp/viewmodels/NowPlayingFragmentViewModel$Factory;", "app_debug"})
public final class InjectorUtils {
    public static final com.example.android.uamp.utils.InjectorUtils INSTANCE = null;
    
    private final com.example.android.uamp.common.MusicServiceConnection provideMusicServiceConnection(android.content.Context context) {
        return null;
    }
    
    @org.jetbrains.annotations.NotNull()
    public final com.example.android.uamp.viewmodels.MainActivityViewModel.Factory provideMainActivityViewModel(@org.jetbrains.annotations.NotNull()
    android.content.Context context) {
        return null;
    }
    
    @org.jetbrains.annotations.NotNull()
    public final com.example.android.uamp.viewmodels.MediaItemFragmentViewModel.Factory provideMediaItemFragmentViewModel(@org.jetbrains.annotations.NotNull()
    android.content.Context context, @org.jetbrains.annotations.NotNull()
    java.lang.String mediaId) {
        return null;
    }
    
    @org.jetbrains.annotations.NotNull()
    public final com.example.android.uamp.viewmodels.NowPlayingFragmentViewModel.Factory provideNowPlayingFragmentViewModel(@org.jetbrains.annotations.NotNull()
    android.content.Context context) {
        return null;
    }
    
    private InjectorUtils() {
        super();
    }
}