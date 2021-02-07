package com.example.libgovskotlin

import android.os.Build
import android.os.Bundle
import android.os.StrictMode
import androidx.appcompat.app.AppCompatActivity
import androidx.navigation.findNavController
import androidx.navigation.ui.AppBarConfiguration
import androidx.navigation.ui.setupActionBarWithNavController
import androidx.navigation.ui.setupWithNavController
import com.google.android.material.bottomnavigation.BottomNavigationView
import gomobile.Gomobile
import java.net.HttpURLConnection
import java.net.URL


class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        print("Kotlin\n")

        println("lol")
        if (Build.VERSION.SDK_INT > 9) {
            val policy = StrictMode.ThreadPolicy.Builder().permitAll().build()
            StrictMode.setThreadPolicy(policy)
        }
        val jsonStr = URL("http://192.168.0.31:5001/ID").readText()
//        val jsonStr = URL("http://0.0.0.0:5001/ID").readText()
        println(jsonStr)

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


        val ret = Gomobile.getID()
        print(ret)

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

        print("Kotlin\n")


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
}