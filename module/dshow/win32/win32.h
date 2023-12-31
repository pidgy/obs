#ifdef __cplusplus
extern "C" {
#endif

#include <stdlib.h>

typedef struct _VideoCaptureDevice {
    char *name;
    char *path;

    int namelen;
    int pathlen;
} VideoCaptureDevice;

int GetVideoCaptureDevice(int index, VideoCaptureDevice *device);

const char * GetVideoCaptureDeviceName(int device, int *len);
const char * GetVideoCaptureDevicePath(int device, int *len);
const char * GetVideoCaptureDeviceDescription(int device, int *len);
const char * GetVideoCaptureDeviceWaveInID(int device, int *len);


#ifdef __cplusplus
}
#endif
