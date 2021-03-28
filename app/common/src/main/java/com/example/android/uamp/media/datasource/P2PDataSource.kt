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
//import kotlinLib.KotlinLib
import poc2PeerKotlinInterface.Poc2PeerKotlinInterface
import java.util.*

/** A [DataSource] for reading P2P urls, via Bittorrent.  */
class P2PDataSource
    : BaseDataSource( /* isNetwork= */true) {
    private var dataSpec: DataSpec? = null
    private var data: ByteArray? = null
    private var endPosition = 0L
    private var readPosition = 0L
    private var ID = ""

    @Throws(IOException::class)
    override fun open(dataSpec: DataSpec): Long {
        transferInitializing(dataSpec)
        Log.d("p2psource", dataSpec.toString())
        this.dataSpec = dataSpec
        readPosition = dataSpec.position
        val uri = dataSpec.uri
        ID = uri.schemeSpecificPart.removePrefix("//")
        endPosition = Poc2PeerKotlinInterface.open(ID)
        transferStarted(dataSpec)
        return endPosition - readPosition
    }

    override fun read(buffer: ByteArray, offset: Int, readLength: Int): Int {
        var readLength = readLength.toLong()
        if (readLength == 0L) {
            return 0
        }
        val remainingBytes = endPosition - readPosition
        if (remainingBytes == 0L) {
            return C.RESULT_END_OF_INPUT
        }
        readLength = readLength.coerceAtMost(remainingBytes)
        val data = Poc2PeerKotlinInterface.read(buffer, readPosition, offset.toLong(), readLength, ID);
        System.arraycopy(Util.castNonNull(data), 0, buffer, offset, readLength.toInt())
        readPosition += data.size
        bytesTransferred(data.size)
        return data.size
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