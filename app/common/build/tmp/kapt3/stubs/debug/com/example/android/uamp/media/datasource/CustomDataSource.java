package com.example.android.uamp.media.datasource;

import java.lang.System;

/**
 * A [DataSource] that supports multiple URI schemes. The supported schemes are:
 *
 *
 * * file: For fetching data from a local file (e.g. file:///path/to/media/media.mp4, or just
 * /path/to/media/media.mp4 because the implementation assumes that a URI without a scheme is
 * a local file URI).
 * * asset: For fetching data from an asset in the application's apk (e.g. asset:///media.mp4).
 * * rawresource: For fetching data from a raw resource in the application's apk (e.g.
 * rawresource:///resourceId, where rawResourceId is the integer identifier of the raw
 * resource).
 * * content: For fetching data from a content URI (e.g. content://authority/path/123).
 * * rtmp: For fetching data over RTMP. Only supported if the project using ExoPlayer has an
 * explicit dependency on ExoPlayer's RTMP extension.
 * * data: For parsing data inlined in the URI as defined in RFC 2397.
 * * udp: For fetching data over UDP (e.g. udp://something.com/media).
 * * http(s): For fetching data over HTTP and HTTPS (e.g. https://www.something.com/media.mp4),
 * if constructed using [.DefaultDataSource], or any other
 * schemes supported by a base data source if constructed using [       ][.DefaultDataSource].
 */
@kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000d\n\u0002\u0018\u0002\n\u0002\u0018\u0002\n\u0000\n\u0002\u0018\u0002\n\u0000\n\u0002\u0010\u000e\n\u0000\n\u0002\u0010\u000b\n\u0002\b\u0002\n\u0002\u0010\b\n\u0002\b\f\n\u0002\u0010!\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0010\u0002\n\u0002\b\t\n\u0002\u0010$\n\u0002\u0010 \n\u0002\b\u0003\n\u0002\u0018\u0002\n\u0002\b\u0003\n\u0002\u0010\t\n\u0000\n\u0002\u0018\u0002\n\u0002\b\u0002\n\u0002\u0010\u0012\n\u0002\b\u0004\u0018\u0000 62\u00020\u0001:\u00016B!\b\u0016\u0012\u0006\u0010\u0002\u001a\u00020\u0003\u0012\b\u0010\u0004\u001a\u0004\u0018\u00010\u0005\u0012\u0006\u0010\u0006\u001a\u00020\u0007\u00a2\u0006\u0002\u0010\bB1\b\u0016\u0012\u0006\u0010\u0002\u001a\u00020\u0003\u0012\b\u0010\u0004\u001a\u0004\u0018\u00010\u0005\u0012\u0006\u0010\t\u001a\u00020\n\u0012\u0006\u0010\u000b\u001a\u00020\n\u0012\u0006\u0010\u0006\u001a\u00020\u0007\u00a2\u0006\u0002\u0010\fB\u0017\u0012\u0006\u0010\u0002\u001a\u00020\u0003\u0012\b\u0010\r\u001a\u0004\u0018\u00010\u0001\u00a2\u0006\u0002\u0010\u000eJ\u0012\u0010\u001a\u001a\u00020\u001b2\b\u0010\u0012\u001a\u0004\u0018\u00010\u0001H\u0002J\u0010\u0010\u001c\u001a\u00020\u001b2\u0006\u0010\u001d\u001a\u00020\u0018H\u0016J\b\u0010\u001e\u001a\u00020\u001bH\u0016J\b\u0010\u001f\u001a\u00020\u0001H\u0002J\b\u0010 \u001a\u00020\u0001H\u0002J\b\u0010!\u001a\u00020\u0001H\u0002J\b\u0010\"\u001a\u00020\u0001H\u0002J\b\u0010#\u001a\u00020\u0001H\u0002J\u001a\u0010$\u001a\u0014\u0012\u0004\u0012\u00020\u0005\u0012\n\u0012\b\u0012\u0004\u0012\u00020\u00050&0%H\u0016J\n\u0010\'\u001a\u0004\u0018\u00010\u0001H\u0002J\b\u0010(\u001a\u00020\u0001H\u0002J\n\u0010)\u001a\u0004\u0018\u00010*H\u0016J\u001a\u0010+\u001a\u00020\u001b2\b\u0010\u0012\u001a\u0004\u0018\u00010\u00012\u0006\u0010,\u001a\u00020\u0018H\u0002J\u0010\u0010-\u001a\u00020.2\u0006\u0010/\u001a\u000200H\u0016J \u00101\u001a\u00020\n2\u0006\u00102\u001a\u0002032\u0006\u00104\u001a\u00020\n2\u0006\u00105\u001a\u00020\nH\u0016R\u0010\u0010\u000f\u001a\u0004\u0018\u00010\u0001X\u0082\u000e\u00a2\u0006\u0002\n\u0000R\u000e\u0010\r\u001a\u00020\u0001X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u0010\u0010\u0010\u001a\u0004\u0018\u00010\u0001X\u0082\u000e\u00a2\u0006\u0002\n\u0000R\u000e\u0010\u0002\u001a\u00020\u0003X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u0010\u0010\u0011\u001a\u0004\u0018\u00010\u0001X\u0082\u000e\u00a2\u0006\u0002\n\u0000R\u0010\u0010\u0012\u001a\u0004\u0018\u00010\u0001X\u0082\u000e\u00a2\u0006\u0002\n\u0000R\u0010\u0010\u0013\u001a\u0004\u0018\u00010\u0001X\u0082\u000e\u00a2\u0006\u0002\n\u0000R\u0010\u0010\u0014\u001a\u0004\u0018\u00010\u0001X\u0082\u000e\u00a2\u0006\u0002\n\u0000R\u0010\u0010\u0015\u001a\u0004\u0018\u00010\u0001X\u0082\u000e\u00a2\u0006\u0002\n\u0000R\u0014\u0010\u0016\u001a\b\u0012\u0004\u0012\u00020\u00180\u0017X\u0082\u0004\u00a2\u0006\u0002\n\u0000R\u0010\u0010\u0019\u001a\u0004\u0018\u00010\u0001X\u0082\u000e\u00a2\u0006\u0002\n\u0000\u00a8\u00067"}, d2 = {"Lcom/example/android/uamp/media/datasource/CustomDataSource;", "Lcom/google/android/exoplayer2/upstream/DataSource;", "context", "Landroid/content/Context;", "userAgent", "", "allowCrossProtocolRedirects", "", "(Landroid/content/Context;Ljava/lang/String;Z)V", "connectTimeoutMillis", "", "readTimeoutMillis", "(Landroid/content/Context;Ljava/lang/String;IIZ)V", "baseDataSource", "(Landroid/content/Context;Lcom/google/android/exoplayer2/upstream/DataSource;)V", "assetDataSource", "contentDataSource", "dataSchemeDataSource", "dataSource", "fileDataSource", "rawResourceDataSource", "rtmpDataSource", "transferListeners", "", "Lcom/google/android/exoplayer2/upstream/TransferListener;", "udpDataSource", "addListenersToDataSource", "", "addTransferListener", "transferListener", "close", "getAssetDataSource", "getContentDataSource", "getDataSchemeDataSource", "getFileDataSource", "getRawResourceDataSource", "getResponseHeaders", "", "", "getRtmpDataSource", "getUdpDataSource", "getUri", "Landroid/net/Uri;", "maybeAddListenerToDataSource", "listener", "open", "", "dataSpec", "Lcom/google/android/exoplayer2/upstream/DataSpec;", "read", "buffer", "", "offset", "readLength", "Companion", "common_debug"})
public final class CustomDataSource implements com.google.android.exoplayer2.upstream.DataSource {
    private final android.content.Context context = null;
    private final java.util.List<com.google.android.exoplayer2.upstream.TransferListener> transferListeners = null;
    private final com.google.android.exoplayer2.upstream.DataSource baseDataSource = null;
    private com.google.android.exoplayer2.upstream.DataSource fileDataSource;
    private com.google.android.exoplayer2.upstream.DataSource assetDataSource;
    private com.google.android.exoplayer2.upstream.DataSource contentDataSource;
    private com.google.android.exoplayer2.upstream.DataSource rtmpDataSource;
    private com.google.android.exoplayer2.upstream.DataSource udpDataSource;
    private com.google.android.exoplayer2.upstream.DataSource dataSchemeDataSource;
    private com.google.android.exoplayer2.upstream.DataSource rawResourceDataSource;
    private com.google.android.exoplayer2.upstream.DataSource dataSource;
    private static final java.lang.String TAG = "DefaultDataSource";
    private static final java.lang.String SCHEME_ASSET = "asset";
    private static final java.lang.String SCHEME_CONTENT = "content";
    private static final java.lang.String SCHEME_RTMP = "rtmp";
    private static final java.lang.String SCHEME_UDP = "udp";
    private static final java.lang.String SCHEME_RAW = "rawresource";
    public static final com.example.android.uamp.media.datasource.CustomDataSource.Companion Companion = null;
    
    @java.lang.Override()
    public void addTransferListener(@org.jetbrains.annotations.NotNull()
    com.google.android.exoplayer2.upstream.TransferListener transferListener) {
    }
    
    @java.lang.Override()
    public long open(@org.jetbrains.annotations.NotNull()
    com.google.android.exoplayer2.upstream.DataSpec dataSpec) throws java.io.IOException {
        return 0L;
    }
    
    @java.lang.Override()
    public int read(@org.jetbrains.annotations.NotNull()
    byte[] buffer, int offset, int readLength) throws java.io.IOException {
        return 0;
    }
    
    @org.jetbrains.annotations.Nullable()
    @java.lang.Override()
    public android.net.Uri getUri() {
        return null;
    }
    
    @org.jetbrains.annotations.NotNull()
    @java.lang.Override()
    public java.util.Map<java.lang.String, java.util.List<java.lang.String>> getResponseHeaders() {
        return null;
    }
    
    @java.lang.Override()
    public void close() throws java.io.IOException {
    }
    
    private final com.google.android.exoplayer2.upstream.DataSource getUdpDataSource() {
        return null;
    }
    
    private final com.google.android.exoplayer2.upstream.DataSource getFileDataSource() {
        return null;
    }
    
    private final com.google.android.exoplayer2.upstream.DataSource getAssetDataSource() {
        return null;
    }
    
    private final com.google.android.exoplayer2.upstream.DataSource getContentDataSource() {
        return null;
    }
    
    private final com.google.android.exoplayer2.upstream.DataSource getRtmpDataSource() {
        return null;
    }
    
    private final com.google.android.exoplayer2.upstream.DataSource getDataSchemeDataSource() {
        return null;
    }
    
    private final com.google.android.exoplayer2.upstream.DataSource getRawResourceDataSource() {
        return null;
    }
    
    private final void addListenersToDataSource(com.google.android.exoplayer2.upstream.DataSource dataSource) {
    }
    
    private final void maybeAddListenerToDataSource(com.google.android.exoplayer2.upstream.DataSource dataSource, com.google.android.exoplayer2.upstream.TransferListener listener) {
    }
    
    public CustomDataSource(@org.jetbrains.annotations.NotNull()
    android.content.Context context, @org.jetbrains.annotations.Nullable()
    com.google.android.exoplayer2.upstream.DataSource baseDataSource) {
        super();
    }
    
    /**
     * Constructs a new instance, optionally configured to follow cross-protocol redirects.
     *
     * @param context A context.
     * @param userAgent The User-Agent to use when requesting remote data.
     * @param allowCrossProtocolRedirects Whether cross-protocol redirects (i.e. redirects from HTTP
     * to HTTPS and vice versa) are enabled when fetching remote data.
     */
    public CustomDataSource(@org.jetbrains.annotations.NotNull()
    android.content.Context context, @org.jetbrains.annotations.Nullable()
    java.lang.String userAgent, boolean allowCrossProtocolRedirects) {
        super();
    }
    
    /**
     * Constructs a new instance, optionally configured to follow cross-protocol redirects.
     *
     * @param context A context.
     * @param userAgent The User-Agent to use when requesting remote data.
     * @param connectTimeoutMillis The connection timeout that should be used when requesting remote
     * data, in milliseconds. A timeout of zero is interpreted as an infinite timeout.
     * @param readTimeoutMillis The read timeout that should be used when requesting remote data, in
     * milliseconds. A timeout of zero is interpreted as an infinite timeout.
     * @param allowCrossProtocolRedirects Whether cross-protocol redirects (i.e. redirects from HTTP
     * to HTTPS and vice versa) are enabled when fetching remote data.
     */
    public CustomDataSource(@org.jetbrains.annotations.NotNull()
    android.content.Context context, @org.jetbrains.annotations.Nullable()
    java.lang.String userAgent, int connectTimeoutMillis, int readTimeoutMillis, boolean allowCrossProtocolRedirects) {
        super();
    }
    
    @kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000\u0014\n\u0002\u0018\u0002\n\u0002\u0010\u0000\n\u0002\b\u0002\n\u0002\u0010\u000e\n\u0002\b\u0006\b\u0086\u0003\u0018\u00002\u00020\u0001B\u0007\b\u0002\u00a2\u0006\u0002\u0010\u0002R\u000e\u0010\u0003\u001a\u00020\u0004X\u0082T\u00a2\u0006\u0002\n\u0000R\u000e\u0010\u0005\u001a\u00020\u0004X\u0082T\u00a2\u0006\u0002\n\u0000R\u000e\u0010\u0006\u001a\u00020\u0004X\u0082T\u00a2\u0006\u0002\n\u0000R\u000e\u0010\u0007\u001a\u00020\u0004X\u0082T\u00a2\u0006\u0002\n\u0000R\u000e\u0010\b\u001a\u00020\u0004X\u0082T\u00a2\u0006\u0002\n\u0000R\u000e\u0010\t\u001a\u00020\u0004X\u0082T\u00a2\u0006\u0002\n\u0000\u00a8\u0006\n"}, d2 = {"Lcom/example/android/uamp/media/datasource/CustomDataSource$Companion;", "", "()V", "SCHEME_ASSET", "", "SCHEME_CONTENT", "SCHEME_RAW", "SCHEME_RTMP", "SCHEME_UDP", "TAG", "common_debug"})
    public static final class Companion {
        
        private Companion() {
            super();
        }
    }
}