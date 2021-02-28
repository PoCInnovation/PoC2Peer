package com.example.android.uamp.viewmodels;

import java.lang.System;

/**
 * [ViewModel] for [MediaItemFragment].
 */
@kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000Z\n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0010\u000e\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0002\u0010 \n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0003\n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0010\u000b\n\u0002\b\u0003\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0010\b\n\u0000\n\u0002\u0010\u0002\n\u0002\b\u0005\u0018\u00002\u00020\u0001:\u0001!B\u0015\u0012\u0006\u0010\u0002\u001a\u00020\u0003\u0012\u0006\u0010\u0004\u001a\u00020\u0005\u00a2\u0006\u0002\u0010\u0006J\u0010\u0010\u001a\u001a\u00020\u001b2\u0006\u0010\u0002\u001a\u00020\u0003H\u0002J\b\u0010\u001c\u001a\u00020\u001dH\u0014J\u001e\u0010\u001e\u001a\b\u0012\u0004\u0012\u00020\n0\t2\u0006\u0010\u001f\u001a\u00020\u00172\u0006\u0010 \u001a\u00020\u0011H\u0002R\u001a\u0010\u0007\u001a\u000e\u0012\n\u0012\b\u0012\u0004\u0012\u00020\n0\t0\bX\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u000e\u0010\u0002\u001a\u00020\u0003X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u001d\u0010\u000b\u001a\u000e\u0012\n\u0012\b\u0012\u0004\u0012\u00020\n0\t0\f\u00a2\u0006\b\n\u0000\u001a\u0004\b\r\u0010\u000eR\u0014\u0010\u000f\u001a\b\u0012\u0004\u0012\u00020\u00110\u0010X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u000e\u0010\u0004\u001a\u00020\u0005X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u001f\u0010\u0012\u001a\u0010\u0012\f\u0012\n \u0014*\u0004\u0018\u00010\u00130\u00130\f\u00a2\u0006\b\n\u0000\u001a\u0004\b\u0015\u0010\u000eR\u0014\u0010\u0016\u001a\b\u0012\u0004\u0012\u00020\u00170\u0010X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u000e\u0010\u0018\u001a\u00020\u0019X\u0082\u0004\u00a2\u0006\u0002\n\u0000\u00a8\u0006\""}, d2 = {"Lcom/example/android/uamp/viewmodels/MediaItemFragmentViewModel;", "Landroidx/lifecycle/ViewModel;", "mediaId", "", "musicServiceConnection", "Lcom/example/android/uamp/common/MusicServiceConnection;", "(Ljava/lang/String;Lcom/example/android/uamp/common/MusicServiceConnection;)V", "_mediaItems", "Landroidx/lifecycle/MutableLiveData;", "", "Lcom/example/android/uamp/MediaItemData;", "mediaItems", "Landroidx/lifecycle/LiveData;", "getMediaItems", "()Landroidx/lifecycle/LiveData;", "mediaMetadataObserver", "Landroidx/lifecycle/Observer;", "Landroid/support/v4/media/MediaMetadataCompat;", "networkError", "", "kotlin.jvm.PlatformType", "getNetworkError", "playbackStateObserver", "Landroid/support/v4/media/session/PlaybackStateCompat;", "subscriptionCallback", "Landroid/support/v4/media/MediaBrowserCompat$SubscriptionCallback;", "getResourceForMediaId", "", "onCleared", "", "updateState", "playbackState", "mediaMetadata", "Factory", "app_debug"})
public final class MediaItemFragmentViewModel extends androidx.lifecycle.ViewModel {
    
    /**
     * Use a backing property so consumers of mediaItems only get a [LiveData] instance so
     * they don't inadvertently modify it.
     */
    private final androidx.lifecycle.MutableLiveData<java.util.List<com.example.android.uamp.MediaItemData>> _mediaItems = null;
    @org.jetbrains.annotations.NotNull()
    private final androidx.lifecycle.LiveData<java.util.List<com.example.android.uamp.MediaItemData>> mediaItems = null;
    
    /**
     * Pass the status of the [MusicServiceConnection.networkFailure] through.
     */
    @org.jetbrains.annotations.NotNull()
    private final androidx.lifecycle.LiveData<java.lang.Boolean> networkError = null;
    private final android.support.v4.media.MediaBrowserCompat.SubscriptionCallback subscriptionCallback = null;
    
    /**
     * When the session's [PlaybackStateCompat] changes, the [mediaItems] need to be updated
     * so the correct [MediaItemData.playbackRes] is displayed on the active item.
     * (i.e.: play/pause button or blank)
     */
    private final androidx.lifecycle.Observer<android.support.v4.media.session.PlaybackStateCompat> playbackStateObserver = null;
    
    /**
     * When the session's [MediaMetadataCompat] changes, the [mediaItems] need to be updated
     * as it means the currently active item has changed. As a result, the new, and potentially
     * old item (if there was one), both need to have their [MediaItemData.playbackRes]
     * changed. (i.e.: play/pause button or blank)
     */
    private final androidx.lifecycle.Observer<android.support.v4.media.MediaMetadataCompat> mediaMetadataObserver = null;
    
    /**
     * Because there's a complex dance between this [ViewModel] and the [MusicServiceConnection]
     * (which is wrapping a [MediaBrowserCompat] object), the usual guidance of using
     * [Transformations] doesn't quite work.
     *
     * Specifically there's three things that are watched that will cause the single piece of
     * [LiveData] exposed from this class to be updated.
     *
     * [subscriptionCallback] (defined above) is called if/when the children of this
     * ViewModel's [mediaId] changes.
     *
     * [MusicServiceConnection.playbackState] changes state based on the playback state of
     * the player, which can change the [MediaItemData.playbackRes]s in the list.
     *
     * [MusicServiceConnection.nowPlaying] changes based on the item that's being played,
     * which can also change the [MediaItemData.playbackRes]s in the list.
     */
    private final com.example.android.uamp.common.MusicServiceConnection musicServiceConnection = null;
    private final java.lang.String mediaId = null;
    
    @org.jetbrains.annotations.NotNull()
    public final androidx.lifecycle.LiveData<java.util.List<com.example.android.uamp.MediaItemData>> getMediaItems() {
        return null;
    }
    
    @org.jetbrains.annotations.NotNull()
    public final androidx.lifecycle.LiveData<java.lang.Boolean> getNetworkError() {
        return null;
    }
    
    /**
     * Since we use [LiveData.observeForever] above (in [musicServiceConnection]), we want
     * to call [LiveData.removeObserver] here to prevent leaking resources when the [ViewModel]
     * is not longer in use.
     *
     * For more details, see the kdoc on [musicServiceConnection] above.
     */
    @java.lang.Override()
    protected void onCleared() {
    }
    
    private final int getResourceForMediaId(java.lang.String mediaId) {
        return 0;
    }
    
    private final java.util.List<com.example.android.uamp.MediaItemData> updateState(android.support.v4.media.session.PlaybackStateCompat playbackState, android.support.v4.media.MediaMetadataCompat mediaMetadata) {
        return null;
    }
    
    public MediaItemFragmentViewModel(@org.jetbrains.annotations.NotNull()
    java.lang.String mediaId, @org.jetbrains.annotations.NotNull()
    com.example.android.uamp.common.MusicServiceConnection musicServiceConnection) {
        super();
    }
    
    @kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000&\n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0010\u000e\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0003\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0002\u0018\u00002\u00020\u0001B\u0015\u0012\u0006\u0010\u0002\u001a\u00020\u0003\u0012\u0006\u0010\u0004\u001a\u00020\u0005\u00a2\u0006\u0002\u0010\u0006J\'\u0010\u0007\u001a\u0002H\b\"\n\b\u0000\u0010\b*\u0004\u0018\u00010\t2\f\u0010\n\u001a\b\u0012\u0004\u0012\u0002H\b0\u000bH\u0016\u00a2\u0006\u0002\u0010\fR\u000e\u0010\u0002\u001a\u00020\u0003X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u000e\u0010\u0004\u001a\u00020\u0005X\u0082\u0004\u00a2\u0006\u0002\n\u0000\u00a8\u0006\r"}, d2 = {"Lcom/example/android/uamp/viewmodels/MediaItemFragmentViewModel$Factory;", "Landroidx/lifecycle/ViewModelProvider$NewInstanceFactory;", "mediaId", "", "musicServiceConnection", "Lcom/example/android/uamp/common/MusicServiceConnection;", "(Ljava/lang/String;Lcom/example/android/uamp/common/MusicServiceConnection;)V", "create", "T", "Landroidx/lifecycle/ViewModel;", "modelClass", "Ljava/lang/Class;", "(Ljava/lang/Class;)Landroidx/lifecycle/ViewModel;", "app_debug"})
    public static final class Factory extends androidx.lifecycle.ViewModelProvider.NewInstanceFactory {
        private final java.lang.String mediaId = null;
        private final com.example.android.uamp.common.MusicServiceConnection musicServiceConnection = null;
        
        @kotlin.Suppress(names = {"unchecked_cast"})
        @java.lang.Override()
        public <T extends androidx.lifecycle.ViewModel>T create(@org.jetbrains.annotations.NotNull()
        java.lang.Class<T> modelClass) {
            return null;
        }
        
        public Factory(@org.jetbrains.annotations.NotNull()
        java.lang.String mediaId, @org.jetbrains.annotations.NotNull()
        com.example.android.uamp.common.MusicServiceConnection musicServiceConnection) {
            super();
        }
    }
}