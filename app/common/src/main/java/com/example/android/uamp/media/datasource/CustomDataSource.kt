package com.example.android.uamp.media.datasource

import android.content.Context
import android.net.Uri
import com.google.android.exoplayer2.upstream.AssetDataSource
import com.google.android.exoplayer2.upstream.ContentDataSource
import com.google.android.exoplayer2.upstream.DataSchemeDataSource
import com.google.android.exoplayer2.upstream.DataSource
import com.google.android.exoplayer2.upstream.DataSpec
import com.google.android.exoplayer2.upstream.DefaultHttpDataSource
import com.google.android.exoplayer2.upstream.FileDataSource
import com.google.android.exoplayer2.upstream.RawResourceDataSource
import com.google.android.exoplayer2.upstream.TransferListener
import com.google.android.exoplayer2.upstream.UdpDataSource
import com.google.android.exoplayer2.util.Assertions
import com.google.android.exoplayer2.util.Log
import com.google.android.exoplayer2.util.Util
import java.io.IOException

/**
 * A [DataSource] that supports multiple URI schemes. The supported schemes are:
 *
 *
 *  * file: For fetching data from a local file (e.g. file:///path/to/media/media.mp4, or just
 * /path/to/media/media.mp4 because the implementation assumes that a URI without a scheme is
 * a local file URI).
 *  * asset: For fetching data from an asset in the application's apk (e.g. asset:///media.mp4).
 *  * rawresource: For fetching data from a raw resource in the application's apk (e.g.
 * rawresource:///resourceId, where rawResourceId is the integer identifier of the raw
 * resource).
 *  * content: For fetching data from a content URI (e.g. content://authority/path/123).
 *  * rtmp: For fetching data over RTMP. Only supported if the project using ExoPlayer has an
 * explicit dependency on ExoPlayer's RTMP extension.
 *  * data: For parsing data inlined in the URI as defined in RFC 2397.
 *  * udp: For fetching data over UDP (e.g. udp://something.com/media).
 *  * http(s): For fetching data over HTTP and HTTPS (e.g. https://www.something.com/media.mp4),
 * if constructed using [.DefaultDataSource], or any other
 * schemes supported by a base data source if constructed using [       ][.DefaultDataSource].
 *
 */
class CustomDataSource(context: Context, baseDataSource: DataSource?) : DataSource {
    private val context: Context
    private val transferListeners: MutableList<TransferListener>
    private val baseDataSource: DataSource

    // Lazily initialized.
    private var fileDataSource: DataSource? = null
    private var assetDataSource: DataSource? = null
    private var contentDataSource: DataSource? = null
    private var rtmpDataSource: DataSource? = null
    private var udpDataSource: DataSource? = null
    private var dataSchemeDataSource: DataSource? = null
    private var rawResourceDataSource: DataSource? = null
    private var dataSource: DataSource? = null

    /**
     * Constructs a new instance, optionally configured to follow cross-protocol redirects.
     *
     * @param context A context.
     * @param userAgent The User-Agent to use when requesting remote data.
     * @param allowCrossProtocolRedirects Whether cross-protocol redirects (i.e. redirects from HTTP
     * to HTTPS and vice versa) are enabled when fetching remote data.
     */
    constructor(context: Context, userAgent: String?, allowCrossProtocolRedirects: Boolean) : this(
        context,
        userAgent,
        DefaultHttpDataSource.DEFAULT_CONNECT_TIMEOUT_MILLIS,
        DefaultHttpDataSource.DEFAULT_READ_TIMEOUT_MILLIS,
        allowCrossProtocolRedirects
    ) {
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
    constructor(
        context: Context,
        userAgent: String?,
        connectTimeoutMillis: Int,
        readTimeoutMillis: Int,
        allowCrossProtocolRedirects: Boolean
    ) : this(
        context,
        DefaultHttpDataSource(
            userAgent,
            connectTimeoutMillis,
            readTimeoutMillis,
            allowCrossProtocolRedirects,  /* defaultRequestProperties= */
            null
        )
    ) {
    }

    override fun addTransferListener(transferListener: TransferListener) {
        baseDataSource.addTransferListener(transferListener)
        transferListeners.add(transferListener)
        maybeAddListenerToDataSource(fileDataSource, transferListener)
        maybeAddListenerToDataSource(assetDataSource, transferListener)
        maybeAddListenerToDataSource(contentDataSource, transferListener)
        maybeAddListenerToDataSource(rtmpDataSource, transferListener)
        maybeAddListenerToDataSource(udpDataSource, transferListener)
        maybeAddListenerToDataSource(dataSchemeDataSource, transferListener)
        maybeAddListenerToDataSource(rawResourceDataSource, transferListener)
    }

    @Throws(IOException::class)
    override fun open(dataSpec: DataSpec): Long {
        Log.d("customdatasource", "OPEN")
        Assertions.checkState(dataSource == null)
        // Choose the correct source for the scheme.
        val scheme = dataSpec.uri.scheme
        dataSource = if (Util.isLocalFileUri(dataSpec.uri)) {
            val uriPath = dataSpec.uri.path
            if (uriPath != null && uriPath.startsWith("/android_asset/")) {
                getAssetDataSource()
            } else {
                getFileDataSource()
            }
        } else if (CustomDataSource.Companion.SCHEME_ASSET == scheme) {
            getAssetDataSource()
        } else if (CustomDataSource.Companion.SCHEME_CONTENT == scheme) {
            getContentDataSource()
        } else if (CustomDataSource.Companion.SCHEME_RTMP == scheme) {
            getRtmpDataSource()
        } else if (CustomDataSource.Companion.SCHEME_UDP == scheme) {
            getUdpDataSource()
        } else if (DataSchemeDataSource.SCHEME_DATA == scheme) {
            getDataSchemeDataSource()
        } else if (CustomDataSource.Companion.SCHEME_RAW == scheme) {
            getRawResourceDataSource()
        } else {
            baseDataSource
        }
        // Open the source and return.
        return dataSource!!.open(dataSpec)
    }

    @Throws(IOException::class)
    override fun read(buffer: ByteArray, offset: Int, readLength: Int): Int {
        Log.e("customdatasource", "read")
        return Assertions.checkNotNull(dataSource).read(buffer, offset, readLength)
    }

    override fun getUri(): Uri? {
        Log.d("customdatasource", dataSource!!.uri.toString())
        return if (dataSource == null) null else dataSource!!.uri
    }

    override fun getResponseHeaders(): Map<String, List<String>> {
        Log.d("customdatasource", dataSource!!.responseHeaders.toString())
        return if (dataSource == null) emptyMap() else dataSource!!.responseHeaders
    }

    @Throws(IOException::class)
    override fun close() {
        Log.d("customdatasource", "close")
        if (dataSource != null) {
            try {
                dataSource!!.close()
            } finally {
                dataSource = null
            }
        }
    }

    private fun getUdpDataSource(): DataSource {
        Log.d("customdatasource", "getUDP")
        if (udpDataSource == null) {
            udpDataSource = UdpDataSource()
            addListenersToDataSource(udpDataSource)
        }
        return udpDataSource as DataSource
    }

    private fun getFileDataSource(): DataSource {
        Log.d("customdatasource", "getFILE")
        if (fileDataSource == null) {
            fileDataSource = FileDataSource()
            addListenersToDataSource(fileDataSource)
        }
        return fileDataSource as DataSource
    }

    private fun getAssetDataSource(): DataSource {
        Log.d("customdatasource", "getASSET")
        if (assetDataSource == null) {
            assetDataSource = AssetDataSource(context)
            addListenersToDataSource(assetDataSource)
        }
        return assetDataSource as DataSource
    }

    private fun getContentDataSource(): DataSource {
        Log.d("customdatasource", "getCONTENT")
        if (contentDataSource == null) {
            contentDataSource = ContentDataSource(context)
            addListenersToDataSource(contentDataSource)
        }
        return contentDataSource as DataSource
    }

    private fun getRtmpDataSource(): DataSource? {
        Log.d("customdatasource", "getRTMP")
        if (rtmpDataSource == null) {
            try {
                // LINT.IfChange
                val clazz = Class.forName("com.google.android.exoplayer2.ext.rtmp.RtmpDataSource")
                rtmpDataSource = clazz.getConstructor().newInstance() as DataSource
                // LINT.ThenChange(../../../../../../../../proguard-rules.txt)
                addListenersToDataSource(rtmpDataSource)
            } catch (e: ClassNotFoundException) {
                // Expected if the app was built without the RTMP extension.
                Log.w(
                    CustomDataSource.Companion.TAG,
                    "Attempting to play RTMP stream without depending on the RTMP extension"
                )
            } catch (e: Exception) {
                // The RTMP extension is present, but instantiation failed.
                throw RuntimeException("Error instantiating RTMP extension", e)
            }
            if (rtmpDataSource == null) {
                rtmpDataSource = baseDataSource
            }
        }
        return rtmpDataSource
    }

    private fun getDataSchemeDataSource(): DataSource {
        Log.d("customdatasource", "getDATASCHEME")
        if (dataSchemeDataSource == null) {
            dataSchemeDataSource = DataSchemeDataSource()
            addListenersToDataSource(dataSchemeDataSource)
        }
        return dataSchemeDataSource as DataSource
    }

    private fun getRawResourceDataSource(): DataSource {
        Log.d("customdatasource", "getRAWRESOURCE")
        if (rawResourceDataSource == null) {
            rawResourceDataSource = RawResourceDataSource(context)
            addListenersToDataSource(rawResourceDataSource)
        }
        return rawResourceDataSource as DataSource
    }

    private fun addListenersToDataSource(dataSource: DataSource?) {
        Log.d("customdatasource", "addListeners")
        for (i in transferListeners.indices) {
            dataSource!!.addTransferListener(transferListeners[i])
        }
    }

    private fun maybeAddListenerToDataSource(
        dataSource: DataSource?, listener: TransferListener
    ) {
        Log.d("customdatasource", "maybeaddListener")
        dataSource?.addTransferListener(listener)
    }

    companion object {
        private const val TAG = "DefaultDataSource"
        private const val SCHEME_ASSET = "asset"
        private const val SCHEME_CONTENT = "content"
        private const val SCHEME_RTMP = "rtmp"
        private const val SCHEME_UDP = "udp"
        private const val SCHEME_RAW = RawResourceDataSource.RAW_RESOURCE_SCHEME
    }

    /**
     * Constructs a new instance that delegates to a provided [DataSource] for URI schemes other
     * than file, asset and content.
     *
     * @param context A context.
     * @param baseDataSource A [DataSource] to use for URI schemes other than file, asset and
     * content. This [DataSource] should normally support at least http(s).
     */
    init {
        this.context = context.applicationContext
        this.baseDataSource = Assertions.checkNotNull(baseDataSource)
        transferListeners = ArrayList()
    }
}