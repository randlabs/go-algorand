#
# Makefile for WiX-based windows Installer
#
# (c) 2021 Rand Labs.
#
# Author: Hernán Di Pietro <hernan@randlabs.io>
#

SOURCENAME=algorand-node
OBJ=$(SOURCENAME)-$(NETWORK).wixobj
MSI=$(SOURCENAME)-$(NETWORK).msi
CANDLE=$(WIX)bin\candle.exe
LIGHT=$(WIX)bin\light.exe
GOBINFOLDER=$(HOME)/go/bin
SVCBINFOLDER=../../windows/algodsvc

all: $(OBJ) $(MSI)

$(OBJ): $(SOURCENAME).wxs
	"$(CANDLE)" $(SOURCENAME).wxs -arch x64 -dGOBINFOLDER=$(GOBINFOLDER) -dSVCBINFOLDER=$(SVCBINFOLDER) -dNETWORK=$(NETWORK) -out $(OBJ)

$(MSI): $(OBJ)
	"$(LIGHT)" $(OBJ) -out $(MSI) -v

clean:
	@rm -f *.wixobj
	@rm -f *.wixpdb
	@rm -f *.msi





