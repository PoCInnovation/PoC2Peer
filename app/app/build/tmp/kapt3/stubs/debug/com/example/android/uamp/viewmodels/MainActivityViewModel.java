package com.example.android.uamp.viewmodels;

import java.lang.System;

/**
 * Small [ViewModel] that watches a [MusicServiceConnection] to become connected
 * and provides the root/initial media ID of the underlying [MediaBrowserCompat].
 */
@kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000L\n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0010\u000e\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0007\n\u0002\u0010\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0004\n\u0002\u0010\u000b\n\u0002\b\u0004\n\u0002\u0018\u0002\n\u0002\b\u0004\u0018\u00002\u00020\u0001:\u0001#B\r\u0012\u0006\u0010\u0002\u001a\u00020\u0003\u00a2\u0006\u0002\u0010\u0004J\u0010\u0010\u0013\u001a\u00020\u00142\u0006\u0010\u0015\u001a\u00020\u0016H\u0002J\u000e\u0010\u0017\u001a\u00020\u00142\u0006\u0010\u0018\u001a\u00020\u0016J\u0018\u0010\u0019\u001a\u00020\u00142\u0006\u0010\u0015\u001a\u00020\u00162\b\b\u0002\u0010\u001a\u001a\u00020\u001bJ\u000e\u0010\u001c\u001a\u00020\u00142\u0006\u0010\u001d\u001a\u00020\nJ$\u0010\u001e\u001a\u00020\u00142\u0006\u0010\u001f\u001a\u00020 2\b\b\u0002\u0010!\u001a\u00020\u001b2\n\b\u0002\u0010\"\u001a\u0004\u0018\u00010\nR\u001a\u0010\u0005\u001a\u000e\u0012\n\u0012\b\u0012\u0004\u0012\u00020\b0\u00070\u0006X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u001a\u0010\t\u001a\u000e\u0012\n\u0012\b\u0012\u0004\u0012\u00020\n0\u00070\u0006X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u000e\u0010\u0002\u001a\u00020\u0003X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u001d\u0010\u000b\u001a\u000e\u0012\n\u0012\b\u0012\u0004\u0012\u00020\b0\u00070\f8F\u00a2\u0006\u0006\u001a\u0004\b\r\u0010\u000eR\u001d\u0010\u000f\u001a\u000e\u0012\n\u0012\b\u0012\u0004\u0012\u00020\n0\u00070\f8F\u00a2\u0006\u0006\u001a\u0004\b\u0010\u0010\u000eR\u0017\u0010\u0011\u001a\b\u0012\u0004\u0012\u00020\n0\f\u00a2\u0006\b\n\u0000\u001a\u0004\b\u0012\u0010\u000e\u00a8\u0006$"}, d2 = {"Lcom/example/android/uamp/viewmodels/MainActivityViewModel;", "Landroidx/lifecycle/ViewModel;", "musicServiceConnection", "Lcom/example/android/uamp/common/MusicServiceConnection;", "(Lcom/example/android/uamp/common/MusicServiceConnection;)V", "_navigateToFragment", "Landroidx/lifecycle/MutableLiveData;", "Lcom/example/android/uamp/utils/Event;", "Lcom/example/android/uamp/viewmodels/FragmentNavigationRequest;", "_navigateToMediaItem", "", "navigateToFragment", "Landroidx/lifecycle/LiveData;", "getNavigateToFragment", "()Landroidx/lifecycle/LiveData;", "navigateToMediaItem", "getNavigateToMediaItem", "rootMediaId", "getRootMediaId", "browseToItem", "", "mediaItem", "Lcom/example/android/uamp/MediaItemData;", "mediaItemClicked", "clickedItem", "playMedia", "pauseAllowed", "", "playMediaId", "mediaId", "showFragment", "fragment", "Landroidx/fragment/app/Fragment;", "backStack", "tag", "Factory", "app_debug"})
public final class MainActivityViewModel extends androidx.lifecycle.ViewModel {
    @org.jetbrains.annotations.NotNull()
    private final androidx.lifecycle.LiveData<java.lang.String> rootMediaId = null;
    private final androidx.lifecycle.MutableLiveData<com.example.android.uamp.utils.Event<java.lang.String>> _navigateToMediaItem = null;
    private final androidx.lifecycle.MutableLiveData<com.example.android.uamp.utils.Event<com.example.android.uamp.viewmodels.FragmentNavigationRequest>> _navigateToFragment = null;
    private final com.example.android.uamp.common.MusicServiceConnection musicServiceConnection = null;
    
    @org.jetbrains.annotations.NotNull()
    public final androidx.lifecycle.LiveData<java.lang.String> getRootMediaId() {
        return null;
    }
    
    @org.jetbrains.annotations.NotNull()
    public final androidx.lifecycle.LiveData<com.example.android.uamp.utils.Event<java.lang.String>> getNavigateToMediaItem() {
        return null;
    }
    
    @org.jetbrains.annotations.NotNull()
    public final androidx.lifecycle.LiveData<com.example.android.uamp.utils.Event<com.example.android.uamp.viewmodels.FragmentNavigationRequest>> getNavigateToFragment() {
        return null;
    }
    
    /**
     * This method takes a [MediaItemData] and routes it depending on whether it's
     * browsable (i.e.: it's the parent media item of a set of other media items,
     * such as an album), or not.
     *
     * If the item is browsable, handle it by sending an event to the Activity to
     * browse to it, otherwise play it.
     */
    public final void mediaItemClicked(@org.jetbrains.annotations.NotNull()
    com.example.android.uamp.MediaItemData clickedItem) {
    }
    
    /**
     * Convenience method used to swap the fragment shown in the main activity
     *
     * @param fragment the fragment to show
     * @param backStack if true, add this transaction to the back stack
     * @param tag the name to use for this fragment in the stack
     */
    public final void showFragment(@org.jetbrains.annotations.NotNull()
    androidx.fragment.app.Fragment fragment, boolean backStack, @org.jetbrains.annotations.Nullable()
    java.lang.String tag) {
    }
    
    /**
     * This posts a browse [Event] that will be handled by the
     * observer in [MainActivity].
     */
    private final void browseToItem(com.example.android.uamp.MediaItemData mediaItem) {
    }
    
    /**
     * This method takes a [MediaItemData] and does one of the following:
     * - If the item is *not* the active item, then play it directly.
     * - If the item *is* the active item, check whether "pause" is a permitted command. If it is,
     *  then pause playback, otherwise send "play" to resume playback.
     */
    public final void playMedia(@org.jetbrains.annotations.NotNull()
    com.example.android.uamp.MediaItemData mediaItem, boolean pauseAllowed) {
    }
    
    public final void playMediaId(@org.jetbrains.annotations.NotNull()
    java.lang.String mediaId) {
    }
    
    public MainActivityViewModel(@org.jetbrains.annotations.NotNull()
    com.example.android.uamp.common.MusicServiceConnection musicServiceConnection) {
        super();
    }
    
    @kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000 \n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0003\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0002\u0018\u00002\u00020\u0001B\r\u0012\u0006\u0010\u0002\u001a\u00020\u0003\u00a2\u0006\u0002\u0010\u0004J\'\u0010\u0005\u001a\u0002H\u0006\"\n\b\u0000\u0010\u0006*\u0004\u0018\u00010\u00072\f\u0010\b\u001a\b\u0012\u0004\u0012\u0002H\u00060\tH\u0016\u00a2\u0006\u0002\u0010\nR\u000e\u0010\u0002\u001a\u00020\u0003X\u0082\u0004\u00a2\u0006\u0002\n\u0000\u00a8\u0006\u000b"}, d2 = {"Lcom/example/android/uamp/viewmodels/MainActivityViewModel$Factory;", "Landroidx/lifecycle/ViewModelProvider$NewInstanceFactory;", "musicServiceConnection", "Lcom/example/android/uamp/common/MusicServiceConnection;", "(Lcom/example/android/uamp/common/MusicServiceConnection;)V", "create", "T", "Landroidx/lifecycle/ViewModel;", "modelClass", "Ljava/lang/Class;", "(Ljava/lang/Class;)Landroidx/lifecycle/ViewModel;", "app_debug"})
    public static final class Factory extends androidx.lifecycle.ViewModelProvider.NewInstanceFactory {
        private final com.example.android.uamp.common.MusicServiceConnection musicServiceConnection = null;
        
        @kotlin.Suppress(names = {"unchecked_cast"})
        @java.lang.Override()
        public <T extends androidx.lifecycle.ViewModel>T create(@org.jetbrains.annotations.NotNull()
        java.lang.Class<T> modelClass) {
            return null;
        }
        
        public Factory(@org.jetbrains.annotations.NotNull()
        com.example.android.uamp.common.MusicServiceConnection musicServiceConnection) {
            super();
        }
    }
}