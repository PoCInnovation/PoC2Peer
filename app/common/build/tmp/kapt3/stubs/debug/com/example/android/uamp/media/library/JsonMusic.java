package com.example.android.uamp.media.library;

import java.lang.System;

/**
 * An individual piece of music included in our JSON catalog.
 * The format from the server is as specified:
 * ```
 *    { "music" : [
 *    { "title" : // Title of the piece of music
 *    "album" : // Album title of the piece of music
 *    "artist" : // Artist of the piece of music
 *    "genre" : // Primary genre of the music
 *    "source" : // Path to the music, which may be relative
 *    "image" : // Path to the art for the music, which may be relative
 *    "trackNumber" : // Track number
 *    "totalTrackCount" : // Track count
 *    "duration" : // Duration of the music in seconds
 *    "site" : // Source of the music, if applicable
 *    }
 *    ]}
 * ```
 *
 * `source` and `image` can be provided in either relative or
 * absolute paths. For example:
 * ``
 *    "source" : "https://www.example.com/music/ode_to_joy.mp3",
 *    "image" : "ode_to_joy.jpg"
 * ``
 *
 * The `source` specifies the full URI to download the piece of music from, but
 * `image` will be fetched relative to the path of the JSON file itself. This means
 * that if the JSON was at "https://www.example.com/json/music.json" then the image would be found
 * at "https://www.example.com/json/ode_to_joy.jpg".
 */
@kotlin.Suppress(names = {"unused"})
@kotlin.Metadata(mv = {1, 1, 16}, bv = {1, 0, 3}, k = 1, d1 = {"\u0000\u001c\n\u0002\u0018\u0002\n\u0002\u0010\u0000\n\u0002\b\u0002\n\u0002\u0010\u000e\n\u0002\b\b\n\u0002\u0010\t\n\u0002\b\u001d\u0018\u00002\u00020\u0001B\u0005\u00a2\u0006\u0002\u0010\u0002R\u001a\u0010\u0003\u001a\u00020\u0004X\u0086\u000e\u00a2\u0006\u000e\n\u0000\u001a\u0004\b\u0005\u0010\u0006\"\u0004\b\u0007\u0010\bR\u001a\u0010\t\u001a\u00020\u0004X\u0086\u000e\u00a2\u0006\u000e\n\u0000\u001a\u0004\b\n\u0010\u0006\"\u0004\b\u000b\u0010\bR\u001a\u0010\f\u001a\u00020\rX\u0086\u000e\u00a2\u0006\u000e\n\u0000\u001a\u0004\b\u000e\u0010\u000f\"\u0004\b\u0010\u0010\u0011R\u001a\u0010\u0012\u001a\u00020\u0004X\u0086\u000e\u00a2\u0006\u000e\n\u0000\u001a\u0004\b\u0013\u0010\u0006\"\u0004\b\u0014\u0010\bR\u001a\u0010\u0015\u001a\u00020\u0004X\u0086\u000e\u00a2\u0006\u000e\n\u0000\u001a\u0004\b\u0016\u0010\u0006\"\u0004\b\u0017\u0010\bR\u001a\u0010\u0018\u001a\u00020\u0004X\u0086\u000e\u00a2\u0006\u000e\n\u0000\u001a\u0004\b\u0019\u0010\u0006\"\u0004\b\u001a\u0010\bR\u001a\u0010\u001b\u001a\u00020\u0004X\u0086\u000e\u00a2\u0006\u000e\n\u0000\u001a\u0004\b\u001c\u0010\u0006\"\u0004\b\u001d\u0010\bR\u001a\u0010\u001e\u001a\u00020\u0004X\u0086\u000e\u00a2\u0006\u000e\n\u0000\u001a\u0004\b\u001f\u0010\u0006\"\u0004\b \u0010\bR\u001a\u0010!\u001a\u00020\u0004X\u0086\u000e\u00a2\u0006\u000e\n\u0000\u001a\u0004\b\"\u0010\u0006\"\u0004\b#\u0010\bR\u001a\u0010$\u001a\u00020\rX\u0086\u000e\u00a2\u0006\u000e\n\u0000\u001a\u0004\b%\u0010\u000f\"\u0004\b&\u0010\u0011R\u001a\u0010\'\u001a\u00020\rX\u0086\u000e\u00a2\u0006\u000e\n\u0000\u001a\u0004\b(\u0010\u000f\"\u0004\b)\u0010\u0011\u00a8\u0006*"}, d2 = {"Lcom/example/android/uamp/media/library/JsonMusic;", "", "()V", "album", "", "getAlbum", "()Ljava/lang/String;", "setAlbum", "(Ljava/lang/String;)V", "artist", "getArtist", "setArtist", "duration", "", "getDuration", "()J", "setDuration", "(J)V", "genre", "getGenre", "setGenre", "id", "getId", "setId", "image", "getImage", "setImage", "site", "getSite", "setSite", "source", "getSource", "setSource", "title", "getTitle", "setTitle", "totalTrackCount", "getTotalTrackCount", "setTotalTrackCount", "trackNumber", "getTrackNumber", "setTrackNumber", "common_debug"})
public final class JsonMusic {
    @org.jetbrains.annotations.NotNull()
    private java.lang.String id = "";
    @org.jetbrains.annotations.NotNull()
    private java.lang.String title = "";
    @org.jetbrains.annotations.NotNull()
    private java.lang.String album = "";
    @org.jetbrains.annotations.NotNull()
    private java.lang.String artist = "";
    @org.jetbrains.annotations.NotNull()
    private java.lang.String genre = "";
    @org.jetbrains.annotations.NotNull()
    private java.lang.String source = "";
    @org.jetbrains.annotations.NotNull()
    private java.lang.String image = "";
    private long trackNumber = 0L;
    private long totalTrackCount = 0L;
    private long duration = -1L;
    @org.jetbrains.annotations.NotNull()
    private java.lang.String site = "";
    
    @org.jetbrains.annotations.NotNull()
    public final java.lang.String getId() {
        return null;
    }
    
    public final void setId(@org.jetbrains.annotations.NotNull()
    java.lang.String p0) {
    }
    
    @org.jetbrains.annotations.NotNull()
    public final java.lang.String getTitle() {
        return null;
    }
    
    public final void setTitle(@org.jetbrains.annotations.NotNull()
    java.lang.String p0) {
    }
    
    @org.jetbrains.annotations.NotNull()
    public final java.lang.String getAlbum() {
        return null;
    }
    
    public final void setAlbum(@org.jetbrains.annotations.NotNull()
    java.lang.String p0) {
    }
    
    @org.jetbrains.annotations.NotNull()
    public final java.lang.String getArtist() {
        return null;
    }
    
    public final void setArtist(@org.jetbrains.annotations.NotNull()
    java.lang.String p0) {
    }
    
    @org.jetbrains.annotations.NotNull()
    public final java.lang.String getGenre() {
        return null;
    }
    
    public final void setGenre(@org.jetbrains.annotations.NotNull()
    java.lang.String p0) {
    }
    
    @org.jetbrains.annotations.NotNull()
    public final java.lang.String getSource() {
        return null;
    }
    
    public final void setSource(@org.jetbrains.annotations.NotNull()
    java.lang.String p0) {
    }
    
    @org.jetbrains.annotations.NotNull()
    public final java.lang.String getImage() {
        return null;
    }
    
    public final void setImage(@org.jetbrains.annotations.NotNull()
    java.lang.String p0) {
    }
    
    public final long getTrackNumber() {
        return 0L;
    }
    
    public final void setTrackNumber(long p0) {
    }
    
    public final long getTotalTrackCount() {
        return 0L;
    }
    
    public final void setTotalTrackCount(long p0) {
    }
    
    public final long getDuration() {
        return 0L;
    }
    
    public final void setDuration(long p0) {
    }
    
    @org.jetbrains.annotations.NotNull()
    public final java.lang.String getSite() {
        return null;
    }
    
    public final void setSite(@org.jetbrains.annotations.NotNull()
    java.lang.String p0) {
    }
    
    public JsonMusic() {
        super();
    }
}