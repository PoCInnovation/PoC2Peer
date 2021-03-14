package com.example.libgovskotlin

import android.os.Bundle
import androidx.appcompat.app.AppCompatActivity
import androidx.navigation.findNavController
import androidx.navigation.ui.AppBarConfiguration
import androidx.navigation.ui.setupActionBarWithNavController
import androidx.navigation.ui.setupWithNavController
import com.google.android.material.bottomnavigation.BottomNavigationView
import gomobile.Gomobile
import java.net.InetAddress
import java.net.NetworkInterface
import java.net.SocketException
import java.nio.file.Paths
import java.util.*


class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        print("Kotlin\n")

        println(getIPAddress(true))

//        if (Build.VERSION.SDK_INT > 9) {
//            val policy = StrictMode.ThreadPolicy.Builder().permitAll().build()
//            StrictMode.setThreadPolicy(policy)
//        }
//        val jsonStr = URL("http://192.168.0.31:5001/ID").readText()
////        val jsonStr = URL("http://0.0.0.0:5001/ID").readText()
//        println(jsonStr)

//        val url = URL("http://www.android.com/")
//        val urlConnection = url.openConnection() as HttpURLConnection
//        try {
//            val `in`: InputStream = BufferedInputStream(urlConnection.inputStream)
//            readStream(`in`)
//        } finally {
//            urlConnection.disconnect()
//        }
//        // ...
//        // Instantiate the RequestQueue.
//        val queue = Volley.newRequestQueue(this)
//        val url = "https://www.google.com"
//
//        // Request a string response from the provided URL.
//        val stringRequest = StringRequest(Request.Method.GET, url,
//                Response.Listener<String> { response ->
//                    // Display the first 500 characters of the response string.
//                    println("Response is: ${response.substring(0, 500)}")
//                },
//                Response.ErrorListener { textView.text = "That didn't work!" })
//
//        // Add the request to the RequestQueue.
//        queue.add(stringRequest)


//        val url = URL("http://0.0.0.0:5001/ID")
//
//        with(url.openConnection() as HttpURLConnection) {
//            requestMethod = "GET"  // optional default is GET
//
//            println("\nSent 'GET' request to URL : $url; Response Code : $responseCode")
//
//            inputStream.bufferedReader().use {
//                println(it)
//            }
//        }


//        val ret = Gomobile.getID("192.168.0.31:5001")
//        println(ret)
        val ret2 = Gomobile.launchP2P( getIPAddress(true),"192.168.0.6", 5000)
        println(ret2)
        val ret = Gomobile.open("30269e6812313d78c89adc1688e1fdd73d76a79cb2951c0818668c0b96558f02")
        val buf = byteArrayOf()
        val ret1 = Gomobile.read(buf, 0, 20, 20, "30269e6812313d78c89adc1688e1fdd73d76a79cb2951c0818668c0b96558f02")
        println(ret1.toString())

//        val str = Gomobile.callString()
//        print(str)
//
//        val bytes = Gomobile.callByteArray()
//        print(String(bytes))
//        print("Conversion bytes -> String : OK")

//        val ints = Gomobile.callIntArray()
//        print(ints)
//        print("Print Int array directly : OK")
//        for (element in ints) {
//            println(element)
//        }
//        print("Print Int array (Each element) : OK")
//
//        val map = Gomobile.callMap()
//        print("Print Int array interface directly : OK")
//        print(map)
//        for (element in ints) {
//            println(element)
//        }
//        print("Print Int array interface (Each element) : OK")
//
//        val intInterface = Gomobile.callInterfaceInt()
//        print("Print Int array interface directly : OK")
//        print(intInterface)
//        for (element in ints) {
//            println(element)
//        }
//        print("Print Int array interface (Each element) : OK")
//
//        val mapInterface = Gomobile.callInterfaceMap()
//        print(mapInterface)
//        print("Print map directly : OK")
//        for (element in mapInterface) {
//            println(element)
//        }
//        print("Print map interface (Each element) : OK")

//        val packageVarStringified = Gomobile.callPackageVariableStringified()
//        print(packageVarStringified)
//        print("Print packageVar Stringifies: OK")

        println("Kotlin\n")
        println(System.getProperty("user.dir"))
        println(Paths.get("").toAbsolutePath().toString())
        println("Kotlin\n")


        super.onCreate(savedInstanceState)
        setContentView(R.layout.activity_main)
        val navView: BottomNavigationView = findViewById(R.id.nav_view)

        val navController = findNavController(R.id.nav_host_fragment)
        // Passing each menu ID as a set of Ids because each
        // menu should be considered as top level destinations.

        val appBarConfiguration = AppBarConfiguration(setOf(
                R.id.navigation_home, R.id.navigation_dashboard, R.id.navigation_notifications))
        setupActionBarWithNavController(navController, appBarConfiguration)
        navView.setupWithNavController(navController)
    }

    fun getIPAddress(useIPv4: Boolean): String {
        try {
            val interfaces: List<NetworkInterface> = Collections.list(NetworkInterface.getNetworkInterfaces())
            for (intf in interfaces) {
                val addrs: List<InetAddress> = Collections.list(intf.inetAddresses)
                for (addr in addrs) {
                    if (!addr.isLoopbackAddress) {
                        val sAddr = addr.hostAddress
                        val isIPv4 = sAddr.indexOf(':') < 0
                        if (useIPv4) {
                            if (isIPv4) return sAddr
                        } else {
                            if (!isIPv4) {
                                val delim = sAddr.indexOf('%')
                                return if (delim < 0) sAddr.toUpperCase(Locale.ROOT) else sAddr.substring(0, delim).toUpperCase(Locale.ROOT)
                            }
                        }
                    }
                }
            }
        } catch (ignored: Exception) {
        }
        return ""
    }

    private fun getIpAddress(): String {
        var ip = ""
        try {
            val enumNetworkInterfaces: Enumeration<NetworkInterface> = NetworkInterface
                    .getNetworkInterfaces()
            while (enumNetworkInterfaces.hasMoreElements()) {
                val networkInterface: NetworkInterface = enumNetworkInterfaces
                        .nextElement()
                val enumInetAddress: Enumeration<InetAddress> = networkInterface
                        .getInetAddresses()
                while (enumInetAddress.hasMoreElements()) {
                    val inetAddress: InetAddress = enumInetAddress.nextElement()
                    if (inetAddress.isSiteLocalAddress()) {
                        ip += inetAddress.getHostAddress()
                    }
                }
            }
        } catch (e: SocketException) {
            // TODO Auto-generated catch block
            e.printStackTrace()
            ip += """Something Wrong! $e""".trimIndent()
        }
        return ip
    }
}