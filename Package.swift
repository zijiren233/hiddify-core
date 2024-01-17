// swift-tools-version: 5.4
// The swift-tools-version declares the minimum version of Swift required to build this package.

import PackageDescription
import Foundation

let version = "v0.11.0"
let baseURL = "https://github.com/hiddify/hiddify-next-core/releases/download/"
let packageURL = baseURL + version + "/hiddify-libcore-ios.xcframework.zip"

let package = Package(
     name: "Libcore",
     platforms: [
         .iOS(.v13) // Minimum platform version
     ],
     products: [
         .library(
             name: "Libcore",
             targets: ["Libcore"]),
     ],
     dependencies: [
         // No dependencies
     ],
     targets: [
         .binaryTarget(
             name: "Libcore",
             url: packageURL,
             checksum: "4c24525948c624167398c6205004d398c59278ce2551193005d2de95168555bd"
             )
     ]
 )