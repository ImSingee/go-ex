# DT

This is a magical data structure set designed to unify the expression of various data types before Go generics were available.

This package defines several groups of interfaces and follows these conventions:

+ If an interface is named Interface, it must implement the AsInterface method
+ Child interfaces must implement all methods of their parent interfaces

+ Number - All numbers
    + Integer - All integers
      + Int - Signed integers
        + Int8
        + Int16
        + Int32
        + Int64
        + Int128
      + UInt - Unsigned integers
        + UInt8
        + UInt16
        + UInt32
        + UInt64
        + UInt128
    + Float - All floating-point numbers
    + BigDecimal (TODO)