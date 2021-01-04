package com.example.android.uamp.fragments;

import java.lang.System;

/**
 * A fragment representing the current media item being played.
 */
@kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000L\n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0002\b\u0005\n\u0002\u0018\u0002\n\u0002\b\u0005\n\u0002\u0018\u0002\n\u0002\b\u0004\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0010\u0002\n\u0002\b\u0003\n\u0002\u0018\u0002\n\u0002\b\u0002\u0018\u0000 \"2\u00020\u0001:\u0001\"B\u0005\u00a2\u0006\u0002\u0010\u0002J&\u0010\u0014\u001a\u0004\u0018\u00010\u00152\u0006\u0010\u0016\u001a\u00020\u00172\b\u0010\u0018\u001a\u0004\u0018\u00010\u00192\b\u0010\u001a\u001a\u0004\u0018\u00010\u001bH\u0016J\u001a\u0010\u001c\u001a\u00020\u001d2\u0006\u0010\u001e\u001a\u00020\u00152\b\u0010\u001a\u001a\u0004\u0018\u00010\u001bH\u0016J\u0018\u0010\u001f\u001a\u00020\u001d2\u0006\u0010\u001e\u001a\u00020\u00152\u0006\u0010 \u001a\u00020!H\u0002R\u001a\u0010\u0003\u001a\u00020\u0004X\u0086.\u00a2\u0006\u000e\n\u0000\u001a\u0004\b\u0005\u0010\u0006\"\u0004\b\u0007\u0010\bR\u001b\u0010\t\u001a\u00020\n8BX\u0082\u0084\u0002\u00a2\u0006\f\n\u0004\b\r\u0010\u000e\u001a\u0004\b\u000b\u0010\fR\u001b\u0010\u000f\u001a\u00020\u00108BX\u0082\u0084\u0002\u00a2\u0006\f\n\u0004\b\u0013\u0010\u000e\u001a\u0004\b\u0011\u0010\u0012\u00a8\u0006#"}, d2 = {"Lcom/example/android/uamp/fragments/NowPlayingFragment;", "Landroidx/fragment/app/Fragment;", "()V", "binding", "Lcom/example/android/uamp/databinding/FragmentNowplayingBinding;", "getBinding", "()Lcom/example/android/uamp/databinding/FragmentNowplayingBinding;", "setBinding", "(Lcom/example/android/uamp/databinding/FragmentNowplayingBinding;)V", "mainActivityViewModel", "Lcom/example/android/uamp/viewmodels/MainActivityViewModel;", "getMainActivityViewModel", "()Lcom/example/android/uamp/viewmodels/MainActivityViewModel;", "mainActivityViewModel$delegate", "Lkotlin/Lazy;", "nowPlayingViewModel", "Lcom/example/android/uamp/viewmodels/NowPlayingFragmentViewModel;", "getNowPlayingViewModel", "()Lcom/example/android/uamp/viewmodels/NowPlayingFragmentViewModel;", "nowPlayingViewModel$delegate", "onCreateView", "Landroid/view/View;", "inflater", "Landroid/view/LayoutInflater;", "container", "Landroid/view/ViewGroup;", "savedInstanceState", "Landroid/os/Bundle;", "onViewCreated", "", "view", "updateUI", "metadata", "Lcom/example/android/uamp/viewmodels/NowPlayingFragmentViewModel$NowPlayingMetadata;", "Companion", "app_debug"})
public final class NowPlayingFragment extends androidx.fragment.app.Fragment {
    private final kotlin.Lazy mainActivityViewModel$delegate = null;
    private final kotlin.Lazy nowPlayingViewModel$delegate = null;
    @org.jetbrains.annotations.NotNull()
    public com.example.android.uamp.databinding.FragmentNowplayingBinding binding;
    public static final com.example.android.uamp.fragments.NowPlayingFragment.Companion Companion = null;
    private java.util.HashMap _$_findViewCache;
    
    private final com.example.android.uamp.viewmodels.MainActivityViewModel getMainActivityViewModel() {
        return null;
    }
    
    private final com.example.android.uamp.viewmodels.NowPlayingFragmentViewModel getNowPlayingViewModel() {
        return null;
    }
    
    @org.jetbrains.annotations.NotNull()
    public final com.example.android.uamp.databinding.FragmentNowplayingBinding getBinding() {
        return null;
    }
    
    public final void setBinding(@org.jetbrains.annotations.NotNull()
    com.example.android.uamp.databinding.FragmentNowplayingBinding p0) {
    }
    
    @org.jetbrains.annotations.Nullable()
    @java.lang.Override()
    public android.view.View onCreateView(@org.jetbrains.annotations.NotNull()
    android.view.LayoutInflater inflater, @org.jetbrains.annotations.Nullable()
    android.view.ViewGroup container, @org.jetbrains.annotations.Nullable()
    android.os.Bundle savedInstanceState) {
        return null;
    }
    
    @java.lang.Override()
    public void onViewCreated(@org.jetbrains.annotations.NotNull()
    android.view.View view, @org.jetbrains.annotations.Nullable()
    android.os.Bundle savedInstanceState) {
    }
    
    /**
     * Internal function used to update all UI elements except for the current item playback
     */
    private final void updateUI(android.view.View view, com.example.android.uamp.viewmodels.NowPlayingFragmentViewModel.NowPlayingMetadata metadata) {
    }
    
    public NowPlayingFragment() {
        super();
    }
    
    @kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000\u0012\n\u0002\u0018\u0002\n\u0002\u0010\u0000\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0000\b\u0086\u0003\u0018\u00002\u00020\u0001B\u0007\b\u0002\u00a2\u0006\u0002\u0010\u0002J\u0006\u0010\u0003\u001a\u00020\u0004\u00a8\u0006\u0005"}, d2 = {"Lcom/example/android/uamp/fragments/NowPlayingFragment$Companion;", "", "()V", "newInstance", "Lcom/example/android/uamp/fragments/NowPlayingFragment;", "app_debug"})
    public static final class Companion {
        
        @org.jetbrains.annotations.NotNull()
        public final com.example.android.uamp.fragments.NowPlayingFragment newInstance() {
            return null;
        }
        
        private Companion() {
            super();
        }
    }
}