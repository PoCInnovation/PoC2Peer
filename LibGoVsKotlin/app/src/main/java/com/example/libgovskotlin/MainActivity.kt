package com.example.libgovskotlin

import android.os.Bundle
import com.google.android.material.bottomnavigation.BottomNavigationView
import androidx.appcompat.app.AppCompatActivity
import androidx.navigation.findNavController
import androidx.navigation.ui.AppBarConfiguration
import androidx.navigation.ui.setupActionBarWithNavController
import androidx.navigation.ui.setupWithNavController
import gomobile.Gomobile


class MainActivity : AppCompatActivity() {

    override fun onCreate(savedInstanceState: Bundle?) {
        print("Kotlin\n")

        val str = Gomobile.callString()
        print(str)

        val bytes = Gomobile.callByteArray()
        print(String(bytes))
        print("Conversion bytes -> String : OK")

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

        val packageVarStringified = Gomobile.callPackageVariableStringified()
        print(packageVarStringified)
        print("Print packageVar Stringifies: OK")

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