Algorand Node Windows Installer
===============================

How to build
------------

* Install Wix  Toolset (https://wixtoolset.org/).  Version 3.11 or higher is required.
* Build the algorand node following instructions at https://developer.algorand.org/tutorials/compile-and-run-the-algorand-node-natively-windows/ until step 3. Make sure `make` ends successfully.
* Build the Windows Service at this repository under directory `windows/algodsvc`. Follow the `README.md` steps there.
* Execute `make` with the `NETWORK` parameter specifying which kind of node you want to build, e.g to create a testnet node installer you must issue:

    ```
    make NETWORK=testnet
    ```

    The build process should take a while.  It will generate a MSI installer named  `algorand-node-testnet.msi`  or similar depending on your chosen Algorand network.


Installation
------------

To launch the installation process just double click the MSI file, accept the license agreement. 


License
-------
License: AGPL v3

Please see the COPYING_FAQ for details about how to apply our license.