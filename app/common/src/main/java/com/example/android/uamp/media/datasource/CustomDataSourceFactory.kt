/*
 * Copyright (C) 2016 The Android Open Source Project
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

import android.content.Context
import android.util.Log
import com.google.android.exoplayer2.upstream.DataSource
import com.google.android.exoplayer2.upstream.DataSource.Factory
import com.google.android.exoplayer2.upstream.DefaultHttpDataSourceFactory
import com.google.android.exoplayer2.upstream.TransferListener

/**
 * A [Factory] that produces [CustomDataSource] instances that delegate to
 * [DefaultHttpDataSource]s for non-file/asset/content URIs.
 */
class CustomDataSourceFactory(
    context: Context,
    listener: TransferListener?,
    baseDataSourceFactory: DataSource.Factory
) : DataSource.Factory {
    private val context: Context
    private val listener: TransferListener?
    private val baseDataSourceFactory: DataSource.Factory
    /**
     * @param context A context.
     * @param userAgent The User-Agent string that should be used.
     * @param listener An optional listener.
     */
    /**
     * @param context A context.
     * @param userAgent The User-Agent string that should be used.
     */
    @JvmOverloads
    constructor(
        context: Context, userAgent: String?, listener: TransferListener? =  /* listener= */null
    ) : this(context, listener, DefaultHttpDataSourceFactory(userAgent, listener)) {
    }

    /**
     * @param context A context.
     * @param baseDataSourceFactory A [Factory] to be used to create a base [DataSource]
     * for [CustomDataSource].
     * @see CustomDataSource.CustomDataSource
     */
    constructor(
        context: Context,
        baseDataSourceFactory: DataSource.Factory
    ) : this(context,  /* listener= */null, baseDataSourceFactory) {
    }

    override fun createDataSource(): CustomDataSource {
        Log.d("datasource", "data source create")
        val dataSource = CustomDataSource(context, baseDataSourceFactory.createDataSource())
        if (listener != null) {
            dataSource.addTransferListener(listener)
        }
        return dataSource
    }

    /**
     * @param context A context.
     * @param listener An optional listener.
     * @param baseDataSourceFactory A [Factory] to be used to create a base [DataSource]
     * for [CustomDataSource].
     * @see CustomDataSource.CustomDataSource
     */
    init {
        this.context = context.applicationContext
        this.listener = listener
        this.baseDataSourceFactory = baseDataSourceFactory
    }
}