package com.example.android.uamp.common;

import java.lang.System;

/**
 * Class that manages a connection to a [MediaBrowserServiceCompat] instance, typically a
 * [MusicService] or one of its subclasses.
 *
 * Typically it's best to construct/inject dependencies either using DI or, as UAMP does,
 * using [InjectorUtils] in the app module. There are a few difficulties for that here:
 * - [MediaBrowserCompat] is a final class, so mocking it directly is difficult.
 * - A [MediaBrowserConnectionCallback] is a parameter into the construction of
 *  a [MediaBrowserCompat], and provides callbacks to this class.
 * - [MediaBrowserCompat.ConnectionCallback.onConnected] is the best place to construct
 *  a [MediaControllerCompat] that will be used to control the [MediaSessionCompat].
 *
 * Because of these reasons, rather than constructing additional classes, this is treated as
 * a black box (which is why there's very little logic here).
 *
 * This is also why the parameters to construct a [MusicServiceConnection] are simple
 * parameters, rather than private properties. They're only required to build the
 * [MediaBrowserConnectionCallback] and [MediaBrowserCompat] objects.
 */
@kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000v\n\u0002\u0018\u0002\n\u0002\u0010\u0000\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0002\u0010\u000b\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0003\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0010\u000e\n\u0002\b\u0003\n\u0002\u0018\u0002\n\u0002\b\u0005\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\u0010\b\n\u0002\u0010\u0002\n\u0002\b\u0003\n\u0002\u0018\u0002\n\u0002\b\u0005\u0018\u0000 .2\u00020\u0001:\u0003./0B\u0015\u0012\u0006\u0010\u0002\u001a\u00020\u0003\u0012\u0006\u0010\u0004\u001a\u00020\u0005\u00a2\u0006\u0002\u0010\u0006J\u0018\u0010!\u001a\u00020\t2\u0006\u0010\"\u001a\u00020\u001a2\b\u0010#\u001a\u0004\u0018\u00010$J4\u0010!\u001a\u00020\t2\u0006\u0010\"\u001a\u00020\u001a2\b\u0010#\u001a\u0004\u0018\u00010$2\u001a\u0010%\u001a\u0016\u0012\u0004\u0012\u00020\'\u0012\u0006\u0012\u0004\u0018\u00010$\u0012\u0004\u0012\u00020(0&J\u0016\u0010)\u001a\u00020(2\u0006\u0010*\u001a\u00020\u001a2\u0006\u0010+\u001a\u00020,J\u0016\u0010-\u001a\u00020(2\u0006\u0010*\u001a\u00020\u001a2\u0006\u0010+\u001a\u00020,R\u0017\u0010\u0007\u001a\b\u0012\u0004\u0012\u00020\t0\b\u00a2\u0006\b\n\u0000\u001a\u0004\b\u0007\u0010\nR\u000e\u0010\u000b\u001a\u00020\fX\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u0012\u0010\r\u001a\u00060\u000eR\u00020\u0000X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u000e\u0010\u000f\u001a\u00020\u0010X\u0082.\u00a2\u0006\u0002\n\u0000R\u0017\u0010\u0011\u001a\b\u0012\u0004\u0012\u00020\t0\b\u00a2\u0006\b\n\u0000\u001a\u0004\b\u0012\u0010\nR\u0017\u0010\u0013\u001a\b\u0012\u0004\u0012\u00020\u00140\b\u00a2\u0006\b\n\u0000\u001a\u0004\b\u0015\u0010\nR\u0017\u0010\u0016\u001a\b\u0012\u0004\u0012\u00020\u00170\b\u00a2\u0006\b\n\u0000\u001a\u0004\b\u0018\u0010\nR\u0011\u0010\u0019\u001a\u00020\u001a8F\u00a2\u0006\u0006\u001a\u0004\b\u001b\u0010\u001cR\u0011\u0010\u001d\u001a\u00020\u001e8F\u00a2\u0006\u0006\u001a\u0004\b\u001f\u0010 \u00a8\u00061"}, d2 = {"Lcom/example/android/uamp/common/MusicServiceConnection;", "", "context", "Landroid/content/Context;", "serviceComponent", "Landroid/content/ComponentName;", "(Landroid/content/Context;Landroid/content/ComponentName;)V", "isConnected", "Landroidx/lifecycle/MutableLiveData;", "", "()Landroidx/lifecycle/MutableLiveData;", "mediaBrowser", "Landroid/support/v4/media/MediaBrowserCompat;", "mediaBrowserConnectionCallback", "Lcom/example/android/uamp/common/MusicServiceConnection$MediaBrowserConnectionCallback;", "mediaController", "Landroid/support/v4/media/session/MediaControllerCompat;", "networkFailure", "getNetworkFailure", "nowPlaying", "Landroid/support/v4/media/MediaMetadataCompat;", "getNowPlaying", "playbackState", "Landroid/support/v4/media/session/PlaybackStateCompat;", "getPlaybackState", "rootMediaId", "", "getRootMediaId", "()Ljava/lang/String;", "transportControls", "Landroid/support/v4/media/session/MediaControllerCompat$TransportControls;", "getTransportControls", "()Landroid/support/v4/media/session/MediaControllerCompat$TransportControls;", "sendCommand", "command", "parameters", "Landroid/os/Bundle;", "resultCallback", "Lkotlin/Function2;", "", "", "subscribe", "parentId", "callback", "Landroid/support/v4/media/MediaBrowserCompat$SubscriptionCallback;", "unsubscribe", "Companion", "MediaBrowserConnectionCallback", "MediaControllerCallback", "common_debug"})
public final class MusicServiceConnection {
    @org.jetbrains.annotations.NotNull()
    private final androidx.lifecycle.MutableLiveData<java.lang.Boolean> isConnected = null;
    @org.jetbrains.annotations.NotNull()
    private final androidx.lifecycle.MutableLiveData<java.lang.Boolean> networkFailure = null;
    @org.jetbrains.annotations.NotNull()
    private final androidx.lifecycle.MutableLiveData<android.support.v4.media.session.PlaybackStateCompat> playbackState = null;
    @org.jetbrains.annotations.NotNull()
    private final androidx.lifecycle.MutableLiveData<android.support.v4.media.MediaMetadataCompat> nowPlaying = null;
    private final com.example.android.uamp.common.MusicServiceConnection.MediaBrowserConnectionCallback mediaBrowserConnectionCallback = null;
    private final android.support.v4.media.MediaBrowserCompat mediaBrowser = null;
    private android.support.v4.media.session.MediaControllerCompat mediaController;
    private static volatile com.example.android.uamp.common.MusicServiceConnection instance;
    public static final com.example.android.uamp.common.MusicServiceConnection.Companion Companion = null;
    
    @org.jetbrains.annotations.NotNull()
    public final androidx.lifecycle.MutableLiveData<java.lang.Boolean> isConnected() {
        return null;
    }
    
    @org.jetbrains.annotations.NotNull()
    public final androidx.lifecycle.MutableLiveData<java.lang.Boolean> getNetworkFailure() {
        return null;
    }
    
    @org.jetbrains.annotations.NotNull()
    public final java.lang.String getRootMediaId() {
        return null;
    }
    
    @org.jetbrains.annotations.NotNull()
    public final androidx.lifecycle.MutableLiveData<android.support.v4.media.session.PlaybackStateCompat> getPlaybackState() {
        return null;
    }
    
    @org.jetbrains.annotations.NotNull()
    public final androidx.lifecycle.MutableLiveData<android.support.v4.media.MediaMetadataCompat> getNowPlaying() {
        return null;
    }
    
    @org.jetbrains.annotations.NotNull()
    public final android.support.v4.media.session.MediaControllerCompat.TransportControls getTransportControls() {
        return null;
    }
    
    public final void subscribe(@org.jetbrains.annotations.NotNull()
    java.lang.String parentId, @org.jetbrains.annotations.NotNull()
    android.support.v4.media.MediaBrowserCompat.SubscriptionCallback callback) {
    }
    
    public final void unsubscribe(@org.jetbrains.annotations.NotNull()
    java.lang.String parentId, @org.jetbrains.annotations.NotNull()
    android.support.v4.media.MediaBrowserCompat.SubscriptionCallback callback) {
    }
    
    public final boolean sendCommand(@org.jetbrains.annotations.NotNull()
    java.lang.String command, @org.jetbrains.annotations.Nullable()
    android.os.Bundle parameters) {
        return false;
    }
    
    public final boolean sendCommand(@org.jetbrains.annotations.NotNull()
    java.lang.String command, @org.jetbrains.annotations.Nullable()
    android.os.Bundle parameters, @org.jetbrains.annotations.NotNull()
    kotlin.jvm.functions.Function2<? super java.lang.Integer, ? super android.os.Bundle, kotlin.Unit> resultCallback) {
        return false;
    }
    
    public MusicServiceConnection(@org.jetbrains.annotations.NotNull()
    android.content.Context context, @org.jetbrains.annotations.NotNull()
    android.content.ComponentName serviceComponent) {
        super();
    }
    
    @kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000\u001a\n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0010\u0002\n\u0002\b\u0003\b\u0082\u0004\u0018\u00002\u00020\u0001B\r\u0012\u0006\u0010\u0002\u001a\u00020\u0003\u00a2\u0006\u0002\u0010\u0004J\b\u0010\u0005\u001a\u00020\u0006H\u0016J\b\u0010\u0007\u001a\u00020\u0006H\u0016J\b\u0010\b\u001a\u00020\u0006H\u0016R\u000e\u0010\u0002\u001a\u00020\u0003X\u0082\u0004\u00a2\u0006\u0002\n\u0000\u00a8\u0006\t"}, d2 = {"Lcom/example/android/uamp/common/MusicServiceConnection$MediaBrowserConnectionCallback;", "Landroid/support/v4/media/MediaBrowserCompat$ConnectionCallback;", "context", "Landroid/content/Context;", "(Lcom/example/android/uamp/common/MusicServiceConnection;Landroid/content/Context;)V", "onConnected", "", "onConnectionFailed", "onConnectionSuspended", "common_debug"})
    final class MediaBrowserConnectionCallback extends android.support.v4.media.MediaBrowserCompat.ConnectionCallback {
        private final android.content.Context context = null;
        
        /**
         * Invoked after [MediaBrowserCompat.connect] when the request has successfully
         * completed.
         */
        @java.lang.Override()
        public void onConnected() {
        }
        
        /**
         * Invoked when the client is disconnected from the media browser.
         */
        @java.lang.Override()
        public void onConnectionSuspended() {
        }
        
        /**
         * Invoked when the connection to the media browser failed.
         */
        @java.lang.Override()
        public void onConnectionFailed() {
        }
        
        public MediaBrowserConnectionCallback(@org.jetbrains.annotations.NotNull()
        android.content.Context context) {
            super();
        }
    }
    
    @kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000:\n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0010\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0010!\n\u0002\u0018\u0002\n\u0002\b\u0003\n\u0002\u0010\u000e\n\u0000\n\u0002\u0018\u0002\n\u0000\b\u0082\u0004\u0018\u00002\u00020\u0001B\u0005\u00a2\u0006\u0002\u0010\u0002J\u0012\u0010\u0003\u001a\u00020\u00042\b\u0010\u0005\u001a\u0004\u0018\u00010\u0006H\u0016J\u0012\u0010\u0007\u001a\u00020\u00042\b\u0010\b\u001a\u0004\u0018\u00010\tH\u0016J\u0018\u0010\n\u001a\u00020\u00042\u000e\u0010\u000b\u001a\n\u0012\u0004\u0012\u00020\r\u0018\u00010\fH\u0016J\b\u0010\u000e\u001a\u00020\u0004H\u0016J\u001c\u0010\u000f\u001a\u00020\u00042\b\u0010\u0010\u001a\u0004\u0018\u00010\u00112\b\u0010\u0012\u001a\u0004\u0018\u00010\u0013H\u0016\u00a8\u0006\u0014"}, d2 = {"Lcom/example/android/uamp/common/MusicServiceConnection$MediaControllerCallback;", "Landroid/support/v4/media/session/MediaControllerCompat$Callback;", "(Lcom/example/android/uamp/common/MusicServiceConnection;)V", "onMetadataChanged", "", "metadata", "Landroid/support/v4/media/MediaMetadataCompat;", "onPlaybackStateChanged", "state", "Landroid/support/v4/media/session/PlaybackStateCompat;", "onQueueChanged", "queue", "", "Landroid/support/v4/media/session/MediaSessionCompat$QueueItem;", "onSessionDestroyed", "onSessionEvent", "event", "", "extras", "Landroid/os/Bundle;", "common_debug"})
    final class MediaControllerCallback extends android.support.v4.media.session.MediaControllerCompat.Callback {
        
        @java.lang.Override()
        public void onPlaybackStateChanged(@org.jetbrains.annotations.Nullable()
        android.support.v4.media.session.PlaybackStateCompat state) {
        }
        
        @java.lang.Override()
        public void onMetadataChanged(@org.jetbrains.annotations.Nullable()
        android.support.v4.media.MediaMetadataCompat metadata) {
        }
        
        @java.lang.Override()
        public void onQueueChanged(@org.jetbrains.annotations.Nullable()
        java.util.List<android.support.v4.media.session.MediaSessionCompat.QueueItem> queue) {
        }
        
        @java.lang.Override()
        public void onSessionEvent(@org.jetbrains.annotations.Nullable()
        java.lang.String event, @org.jetbrains.annotations.Nullable()
        android.os.Bundle extras) {
        }
        
        /**
         * Normally if a [MediaBrowserServiceCompat] drops its connection the callback comes via
         * [MediaControllerCompat.Callback] (here). But since other connection status events
         * are sent to [MediaBrowserCompat.ConnectionCallback], we catch the disconnect here and
         * send it on to the other callback.
         */
        @java.lang.Override()
        public void onSessionDestroyed() {
        }
        
        public MediaControllerCallback() {
            super();
        }
    }
    
    @kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000 \n\u0002\u0018\u0002\n\u0002\u0010\u0000\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\b\u0086\u0003\u0018\u00002\u00020\u0001B\u0007\b\u0002\u00a2\u0006\u0002\u0010\u0002J\u0016\u0010\u0005\u001a\u00020\u00042\u0006\u0010\u0006\u001a\u00020\u00072\u0006\u0010\b\u001a\u00020\tR\u0010\u0010\u0003\u001a\u0004\u0018\u00010\u0004X\u0082\u000e\u00a2\u0006\u0002\n\u0000\u00a8\u0006\n"}, d2 = {"Lcom/example/android/uamp/common/MusicServiceConnection$Companion;", "", "()V", "instance", "Lcom/example/android/uamp/common/MusicServiceConnection;", "getInstance", "context", "Landroid/content/Context;", "serviceComponent", "Landroid/content/ComponentName;", "common_debug"})
    public static final class Companion {
        
        @org.jetbrains.annotations.NotNull()
        public final com.example.android.uamp.common.MusicServiceConnection getInstance(@org.jetbrains.annotations.NotNull()
        android.content.Context context, @org.jetbrains.annotations.NotNull()
        android.content.ComponentName serviceComponent) {
            return null;
        }
        
        private Companion() {
            super();
        }
    }
}