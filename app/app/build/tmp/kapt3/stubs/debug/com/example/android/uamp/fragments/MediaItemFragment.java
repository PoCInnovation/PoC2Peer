package com.example.android.uamp.fragments;

import java.lang.System;

/**
 * A fragment representing a list of MediaItems.
 */
@kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000N\n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0005\n\u0002\u0010\u000e\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0004\n\u0002\u0010\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0002\u0018\u0000 \u001e2\u00020\u0001:\u0001\u001eB\u0005\u00a2\u0006\u0002\u0010\u0002J\u0012\u0010\u0014\u001a\u00020\u00152\b\u0010\u0016\u001a\u0004\u0018\u00010\u0017H\u0016J&\u0010\u0018\u001a\u0004\u0018\u00010\u00192\u0006\u0010\u001a\u001a\u00020\u001b2\b\u0010\u001c\u001a\u0004\u0018\u00010\u001d2\b\u0010\u0016\u001a\u0004\u0018\u00010\u0017H\u0016R\u000e\u0010\u0003\u001a\u00020\u0004X\u0082.\u00a2\u0006\u0002\n\u0000R\u000e\u0010\u0005\u001a\u00020\u0006X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u001b\u0010\u0007\u001a\u00020\b8BX\u0082\u0084\u0002\u00a2\u0006\f\n\u0004\b\u000b\u0010\f\u001a\u0004\b\t\u0010\nR\u000e\u0010\r\u001a\u00020\u000eX\u0082.\u00a2\u0006\u0002\n\u0000R\u001b\u0010\u000f\u001a\u00020\u00108BX\u0082\u0084\u0002\u00a2\u0006\f\n\u0004\b\u0013\u0010\f\u001a\u0004\b\u0011\u0010\u0012\u00a8\u0006\u001f"}, d2 = {"Lcom/example/android/uamp/fragments/MediaItemFragment;", "Landroidx/fragment/app/Fragment;", "()V", "binding", "Lcom/example/android/uamp/databinding/FragmentMediaitemListBinding;", "listAdapter", "Lcom/example/android/uamp/MediaItemAdapter;", "mainActivityViewModel", "Lcom/example/android/uamp/viewmodels/MainActivityViewModel;", "getMainActivityViewModel", "()Lcom/example/android/uamp/viewmodels/MainActivityViewModel;", "mainActivityViewModel$delegate", "Lkotlin/Lazy;", "mediaId", "", "mediaItemFragmentViewModel", "Lcom/example/android/uamp/viewmodels/MediaItemFragmentViewModel;", "getMediaItemFragmentViewModel", "()Lcom/example/android/uamp/viewmodels/MediaItemFragmentViewModel;", "mediaItemFragmentViewModel$delegate", "onActivityCreated", "", "savedInstanceState", "Landroid/os/Bundle;", "onCreateView", "Landroid/view/View;", "inflater", "Landroid/view/LayoutInflater;", "container", "Landroid/view/ViewGroup;", "Companion", "app_debug"})
public final class MediaItemFragment extends androidx.fragment.app.Fragment {
    private final kotlin.Lazy mainActivityViewModel$delegate = null;
    private final kotlin.Lazy mediaItemFragmentViewModel$delegate = null;
    private java.lang.String mediaId;
    private com.example.android.uamp.databinding.FragmentMediaitemListBinding binding;
    private final com.example.android.uamp.MediaItemAdapter listAdapter = null;
    public static final com.example.android.uamp.fragments.MediaItemFragment.Companion Companion = null;
    private java.util.HashMap _$_findViewCache;
    
    private final com.example.android.uamp.viewmodels.MainActivityViewModel getMainActivityViewModel() {
        return null;
    }
    
    private final com.example.android.uamp.viewmodels.MediaItemFragmentViewModel getMediaItemFragmentViewModel() {
        return null;
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
    public void onActivityCreated(@org.jetbrains.annotations.Nullable()
    android.os.Bundle savedInstanceState) {
    }
    
    public MediaItemFragment() {
        super();
    }
    
    @kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000\u0018\n\u0002\u0018\u0002\n\u0002\u0010\u0000\n\u0002\b\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0010\u000e\n\u0000\b\u0086\u0003\u0018\u00002\u00020\u0001B\u0007\b\u0002\u00a2\u0006\u0002\u0010\u0002J\u000e\u0010\u0003\u001a\u00020\u00042\u0006\u0010\u0005\u001a\u00020\u0006\u00a8\u0006\u0007"}, d2 = {"Lcom/example/android/uamp/fragments/MediaItemFragment$Companion;", "", "()V", "newInstance", "Lcom/example/android/uamp/fragments/MediaItemFragment;", "mediaId", "", "app_debug"})
    public static final class Companion {
        
        @org.jetbrains.annotations.NotNull()
        public final com.example.android.uamp.fragments.MediaItemFragment newInstance(@org.jetbrains.annotations.NotNull()
        java.lang.String mediaId) {
            return null;
        }
        
        private Companion() {
            super();
        }
    }
}