/*
 * Copyright (C) 2017 The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */
package com.example.android.uamp.media.datasource

import android.net.Uri
import android.util.Base64
import android.util.Log
import com.google.android.exoplayer2.C
import com.google.android.exoplayer2.ParserException
import com.google.android.exoplayer2.util.Util
import com.google.android.exoplayer2.upstream.BaseDataSource
import com.google.android.exoplayer2.upstream.DataSourceException
import com.google.android.exoplayer2.upstream.DataSpec
import java.io.IOException
import java.net.URLDecoder

/** A [DataSource] for reading P2P urls, via Bittorrent.  */
class P2PDataSource
    : BaseDataSource( /* isNetwork= */true) {
    private var dataSpec: DataSpec? = null
    private var data: ByteArray? = null
    private var endPosition = 0
    private var readPosition = 0

    @Throws(IOException::class)
    override fun open(dataSpec: DataSpec): Long {
        transferInitializing(dataSpec)
        Log.d("p2psource", dataSpec.toString())
        this.dataSpec = dataSpec
        readPosition = dataSpec.position.toInt()
        val uri = dataSpec.uri
        val uriParts = Util.split(uri.schemeSpecificPart, ",")
        if (uriParts.size != 2) {
            throw ParserException("Unexpected URI format: $uri")
        }
        val dataString = uriParts[1]
        data = if (uriParts[0].contains(";base64")) {
            try {
                Base64.decode(dataString, 0)
            } catch (e: IllegalArgumentException) {
                throw ParserException("Error while parsing Base64 encoded string: $dataString", e)
            }
        } else {
            // TODO: Add support for other charsets.
            Util.getUtf8Bytes(URLDecoder.decode(dataString, C.ASCII_NAME))
        }
        endPosition =
            if (dataSpec.length != C.LENGTH_UNSET.toLong()) dataSpec.length.toInt() + readPosition else data!!.size
        if (endPosition > data!!.size || readPosition > endPosition) {
            data = null
            throw DataSourceException(DataSourceException.POSITION_OUT_OF_RANGE)
        }
        transferStarted(dataSpec)
        return endPosition.toLong() - readPosition
    }

    override fun read(buffer: ByteArray, offset: Int, readLength: Int): Int {
        var readLength = readLength
        if (readLength == 0) {
            return 0
        }
        val remainingBytes = endPosition - readPosition
        if (remainingBytes == 0) {
            return C.RESULT_END_OF_INPUT
        }
        readLength = Math.min(readLength, remainingBytes)
        System.arraycopy(Util.castNonNull(data), readPosition, buffer, offset, readLength)
        readPosition += readLength
        bytesTransferred(readLength)
        return readLength
    }

    override fun getUri(): Uri? {
        return if (dataSpec != null) dataSpec!!.uri else null
    }

    override fun close() {
        if (data != null) {
            data = null
            transferEnded()
        }
        dataSpec = null
    }
}