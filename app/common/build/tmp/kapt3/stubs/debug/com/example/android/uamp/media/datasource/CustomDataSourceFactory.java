package com.example.android.uamp.media.datasource;

import java.lang.System;

/**
 * A [Factory] that produces [CustomDataSource] instances that delegate to
 * [DefaultHttpDataSource]s for non-file/asset/content URIs.
 */
@kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000$\n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0010\u000e\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0005\n\u0002\u0018\u0002\n\u0000\u0018\u00002\u00020\u0001B%\b\u0017\u0012\u0006\u0010\u0002\u001a\u00020\u0003\u0012\b\u0010\u0004\u001a\u0004\u0018\u00010\u0005\u0012\n\b\u0002\u0010\u0006\u001a\u0004\u0018\u00010\u0007\u00a2\u0006\u0002\u0010\bB\u0017\b\u0016\u0012\u0006\u0010\u0002\u001a\u00020\u0003\u0012\u0006\u0010\t\u001a\u00020\u0001\u00a2\u0006\u0002\u0010\nB\u001f\u0012\u0006\u0010\u0002\u001a\u00020\u0003\u0012\b\u0010\u0006\u001a\u0004\u0018\u00010\u0007\u0012\u0006\u0010\t\u001a\u00020\u0001\u00a2\u0006\u0002\u0010\u000bJ\b\u0010\f\u001a\u00020\rH\u0016R\u000e\u0010\t\u001a\u00020\u0001X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u000e\u0010\u0002\u001a\u00020\u0003X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u0010\u0010\u0006\u001a\u0004\u0018\u00010\u0007X\u0082\u0004\u00a2\u0006\u0002\n\u0000\u00a8\u0006\u000e"}, d2 = {"Lcom/example/android/uamp/media/datasource/CustomDataSourceFactory;", "Lcom/google/android/exoplayer2/upstream/DataSource$Factory;", "context", "Landroid/content/Context;", "userAgent", "", "listener", "Lcom/google/android/exoplayer2/upstream/TransferListener;", "(Landroid/content/Context;Ljava/lang/String;Lcom/google/android/exoplayer2/upstream/TransferListener;)V", "baseDataSourceFactory", "(Landroid/content/Context;Lcom/google/android/exoplayer2/upstream/DataSource$Factory;)V", "(Landroid/content/Context;Lcom/google/android/exoplayer2/upstream/TransferListener;Lcom/google/android/exoplayer2/upstream/DataSource$Factory;)V", "createDataSource", "Lcom/example/android/uamp/media/datasource/CustomDataSource;", "common_debug"})
public final class CustomDataSourceFactory implements com.google.android.exoplayer2.upstream.DataSource.Factory {
    private final android.content.Context context = null;
    private final com.google.android.exoplayer2.upstream.TransferListener listener = null;
    private final com.google.android.exoplayer2.upstream.DataSource.Factory baseDataSourceFactory = null;
    
    @org.jetbrains.annotations.NotNull()
    @java.lang.Override()
    public com.example.android.uamp.media.datasource.CustomDataSource createDataSource() {
        return null;
    }
    
    public CustomDataSourceFactory(@org.jetbrains.annotations.NotNull()
    android.content.Context context, @org.jetbrains.annotations.Nullable()
    com.google.android.exoplayer2.upstream.TransferListener listener, @org.jetbrains.annotations.NotNull()
    com.google.android.exoplayer2.upstream.DataSource.Factory baseDataSourceFactory) {
        super();
    }
    
    /**
     * @param context A context.
     * @param userAgent The User-Agent string that should be used.
     */
    public CustomDataSourceFactory(@org.jetbrains.annotations.NotNull()
    android.content.Context context, @org.jetbrains.annotations.Nullable()
    java.lang.String userAgent, @org.jetbrains.annotations.Nullable()
    com.google.android.exoplayer2.upstream.TransferListener listener) {
        super();
    }
    
    /**
     * @param context A context.
     * @param userAgent The User-Agent string that should be used.
     */
    public CustomDataSourceFactory(@org.jetbrains.annotations.NotNull()
    android.content.Context context, @org.jetbrains.annotations.Nullable()
    java.lang.String userAgent) {
        super();
    }
    
    /**
     * @param context A context.
     * @param baseDataSourceFactory A [Factory] to be used to create a base [DataSource]
     * for [CustomDataSource].
     * @see CustomDataSource.CustomDataSource
     */
    public CustomDataSourceFactory(@org.jetbrains.annotations.NotNull()
    android.content.Context context, @org.jetbrains.annotations.NotNull()
    com.google.android.exoplayer2.upstream.DataSource.Factory baseDataSourceFactory) {
        super();
    }
}