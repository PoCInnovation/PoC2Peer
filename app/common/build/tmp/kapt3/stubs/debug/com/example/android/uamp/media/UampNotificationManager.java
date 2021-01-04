package com.example.android.uamp.media;

import java.lang.System;

/**
 * A wrapper class for ExoPlayer's PlayerNotificationManager. It sets up the notification shown to
 * the user during audio playback and provides track metadata, such as track title and icon image.
 */
@kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000D\n\u0002\u0018\u0002\n\u0002\u0010\u0000\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0010\u0002\n\u0002\b\u0003\u0018\u00002\u00020\u0001:\u0001\u0016B\u001d\u0012\u0006\u0010\u0002\u001a\u00020\u0003\u0012\u0006\u0010\u0004\u001a\u00020\u0005\u0012\u0006\u0010\u0006\u001a\u00020\u0007\u00a2\u0006\u0002\u0010\bJ\u0006\u0010\u0013\u001a\u00020\u0014J\u000e\u0010\u0015\u001a\u00020\u00142\u0006\u0010\r\u001a\u00020\u000eR\u000e\u0010\u0002\u001a\u00020\u0003X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u000e\u0010\t\u001a\u00020\nX\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u000e\u0010\u000b\u001a\u00020\fX\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u0010\u0010\r\u001a\u0004\u0018\u00010\u000eX\u0082\u000e\u00a2\u0006\u0002\n\u0000R\u000e\u0010\u000f\u001a\u00020\u0010X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u000e\u0010\u0011\u001a\u00020\u0012X\u0082\u0004\u00a2\u0006\u0002\n\u0000\u0082\u0002\u0004\n\u0002\b\u0019\u00a8\u0006\u0017"}, d2 = {"Lcom/example/android/uamp/media/UampNotificationManager;", "", "context", "Landroid/content/Context;", "sessionToken", "Landroid/support/v4/media/session/MediaSessionCompat$Token;", "notificationListener", "Lcom/google/android/exoplayer2/ui/PlayerNotificationManager$NotificationListener;", "(Landroid/content/Context;Landroid/support/v4/media/session/MediaSessionCompat$Token;Lcom/google/android/exoplayer2/ui/PlayerNotificationManager$NotificationListener;)V", "notificationManager", "Lcom/google/android/exoplayer2/ui/PlayerNotificationManager;", "platformNotificationManager", "Landroid/app/NotificationManager;", "player", "Lcom/google/android/exoplayer2/Player;", "serviceJob", "Lkotlinx/coroutines/Job;", "serviceScope", "Lkotlinx/coroutines/CoroutineScope;", "hideNotification", "", "showNotificationForPlayer", "DescriptionAdapter", "common_debug"})
public final class UampNotificationManager {
    private com.google.android.exoplayer2.Player player;
    private final kotlinx.coroutines.Job serviceJob = null;
    private final kotlinx.coroutines.CoroutineScope serviceScope = null;
    private final com.google.android.exoplayer2.ui.PlayerNotificationManager notificationManager = null;
    private final android.app.NotificationManager platformNotificationManager = null;
    private final android.content.Context context = null;
    
    public final void hideNotification() {
    }
    
    public final void showNotificationForPlayer(@org.jetbrains.annotations.NotNull()
    com.google.android.exoplayer2.Player player) {
    }
    
    public UampNotificationManager(@org.jetbrains.annotations.NotNull()
    android.content.Context context, @org.jetbrains.annotations.NotNull()
    android.support.v4.media.session.MediaSessionCompat.Token sessionToken, @org.jetbrains.annotations.NotNull()
    com.google.android.exoplayer2.ui.PlayerNotificationManager.NotificationListener notificationListener) {
        super();
    }
    
    @kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000B\n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0002\b\u0005\n\u0002\u0018\u0002\n\u0002\b\u0005\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0010\u000e\n\u0002\b\u0003\n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0002\b\u0004\b\u0082\u0004\u0018\u00002\u00020\u0001B\r\u0012\u0006\u0010\u0002\u001a\u00020\u0003\u00a2\u0006\u0002\u0010\u0004J\u0012\u0010\u0011\u001a\u0004\u0018\u00010\u00122\u0006\u0010\u0013\u001a\u00020\u0014H\u0016J\u0010\u0010\u0015\u001a\u00020\u00162\u0006\u0010\u0013\u001a\u00020\u0014H\u0016J\u0010\u0010\u0017\u001a\u00020\u00162\u0006\u0010\u0013\u001a\u00020\u0014H\u0016J\u001e\u0010\u0018\u001a\u0004\u0018\u00010\u00062\u0006\u0010\u0013\u001a\u00020\u00142\n\u0010\u0019\u001a\u00060\u001aR\u00020\u001bH\u0016J\u001b\u0010\u001c\u001a\u0004\u0018\u00010\u00062\u0006\u0010\u001d\u001a\u00020\fH\u0082@\u00f8\u0001\u0000\u00a2\u0006\u0002\u0010\u001eR\u000e\u0010\u0002\u001a\u00020\u0003X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u001c\u0010\u0005\u001a\u0004\u0018\u00010\u0006X\u0086\u000e\u00a2\u0006\u000e\n\u0000\u001a\u0004\b\u0007\u0010\b\"\u0004\b\t\u0010\nR\u001c\u0010\u000b\u001a\u0004\u0018\u00010\fX\u0086\u000e\u00a2\u0006\u000e\n\u0000\u001a\u0004\b\r\u0010\u000e\"\u0004\b\u000f\u0010\u0010\u0082\u0002\u0004\n\u0002\b\u0019\u00a8\u0006\u001f"}, d2 = {"Lcom/example/android/uamp/media/UampNotificationManager$DescriptionAdapter;", "Lcom/google/android/exoplayer2/ui/PlayerNotificationManager$MediaDescriptionAdapter;", "controller", "Landroid/support/v4/media/session/MediaControllerCompat;", "(Lcom/example/android/uamp/media/UampNotificationManager;Landroid/support/v4/media/session/MediaControllerCompat;)V", "currentBitmap", "Landroid/graphics/Bitmap;", "getCurrentBitmap", "()Landroid/graphics/Bitmap;", "setCurrentBitmap", "(Landroid/graphics/Bitmap;)V", "currentIconUri", "Landroid/net/Uri;", "getCurrentIconUri", "()Landroid/net/Uri;", "setCurrentIconUri", "(Landroid/net/Uri;)V", "createCurrentContentIntent", "Landroid/app/PendingIntent;", "player", "Lcom/google/android/exoplayer2/Player;", "getCurrentContentText", "", "getCurrentContentTitle", "getCurrentLargeIcon", "callback", "Lcom/google/android/exoplayer2/ui/PlayerNotificationManager$BitmapCallback;", "Lcom/google/android/exoplayer2/ui/PlayerNotificationManager;", "resolveUriAsBitmap", "uri", "(Landroid/net/Uri;Lkotlin/coroutines/Continuation;)Ljava/lang/Object;", "common_debug"})
    final class DescriptionAdapter implements com.google.android.exoplayer2.ui.PlayerNotificationManager.MediaDescriptionAdapter {
        @org.jetbrains.annotations.Nullable()
        private android.net.Uri currentIconUri;
        @org.jetbrains.annotations.Nullable()
        private android.graphics.Bitmap currentBitmap;
        private final android.support.v4.media.session.MediaControllerCompat controller = null;
        
        @org.jetbrains.annotations.Nullable()
        public final android.net.Uri getCurrentIconUri() {
            return null;
        }
        
        public final void setCurrentIconUri(@org.jetbrains.annotations.Nullable()
        android.net.Uri p0) {
        }
        
        @org.jetbrains.annotations.Nullable()
        public final android.graphics.Bitmap getCurrentBitmap() {
            return null;
        }
        
        public final void setCurrentBitmap(@org.jetbrains.annotations.Nullable()
        android.graphics.Bitmap p0) {
        }
        
        @org.jetbrains.annotations.Nullable()
        @java.lang.Override()
        public android.app.PendingIntent createCurrentContentIntent(@org.jetbrains.annotations.NotNull()
        com.google.android.exoplayer2.Player player) {
            return null;
        }
        
        @org.jetbrains.annotations.NotNull()
        @java.lang.Override()
        public java.lang.String getCurrentContentText(@org.jetbrains.annotations.NotNull()
        com.google.android.exoplayer2.Player player) {
            return null;
        }
        
        @org.jetbrains.annotations.NotNull()
        @java.lang.Override()
        public java.lang.String getCurrentContentTitle(@org.jetbrains.annotations.NotNull()
        com.google.android.exoplayer2.Player player) {
            return null;
        }
        
        @org.jetbrains.annotations.Nullable()
        @java.lang.Override()
        public android.graphics.Bitmap getCurrentLargeIcon(@org.jetbrains.annotations.NotNull()
        com.google.android.exoplayer2.Player player, @org.jetbrains.annotations.NotNull()
        com.google.android.exoplayer2.ui.PlayerNotificationManager.BitmapCallback callback) {
            return null;
        }
        
        public DescriptionAdapter(@org.jetbrains.annotations.NotNull()
        android.support.v4.media.session.MediaControllerCompat controller) {
            super();
        }
    }
}