import React from "react";
import { Link, useLocation } from "react-router-dom";

const Sidebar = () => {
  let location = useLocation().pathname;
console.log(location);

  const routes = [
    {
      title: "Home",
      path: "/",
      icon: (
        <svg
          width="32"
          height="32"
          viewBox="0 0 32 32"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >

            <path
              d="M28 26.6667C28 27.0203 27.8595 27.3594 27.6095 27.6095C27.3594 27.8595 27.0203 28 26.6667 28H5.33333C4.97971 28 4.64057 27.8595 4.39052 27.6095C4.14048 27.3594 4 27.0203 4 26.6667V12.6533C3.99986 12.4502 4.04616 12.2496 4.13535 12.0671C4.22455 11.8845 4.35429 11.7248 4.51467 11.6L15.1813 3.304C15.4154 3.12193 15.7035 3.02308 16 3.02308C16.2965 3.02308 16.5846 3.12193 16.8187 3.304L27.4853 11.6C27.6457 11.7248 27.7754 11.8845 27.8646 12.0671C27.9538 12.2496 28.0001 12.4502 28 12.6533V26.6667ZM25.3333 25.3333V13.304L16 6.04533L6.66667 13.304V25.3333H25.3333Z"
              fill="white"
            />

        </svg>
      ),
    },
    {
      title: "Register new bill",
      path: "/new",
      icon: (
        <svg
          width="32"
          height="32"
          viewBox="0 0 32 32"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
  
            <path
              d="M16 2.66667C23.364 2.66667 29.3333 8.636 29.3333 16C29.3333 23.364 23.364 29.3333 16 29.3333C8.63599 29.3333 2.66666 23.364 2.66666 16C2.66666 8.636 8.63599 2.66667 16 2.66667ZM16 5.33333C10.1093 5.33333 5.33332 10.1093 5.33332 16C5.33332 21.8907 10.1093 26.6667 16 26.6667C21.8907 26.6667 26.6667 21.8907 26.6667 16C26.6667 10.1093 21.8907 5.33333 16 5.33333ZM16 6.66667C17.3573 6.66667 18.6467 6.956 19.8107 7.47733L17.7267 9.56C17.176 9.41333 16.5973 9.33333 16 9.33333C12.3187 9.33333 9.33332 12.3187 9.33332 16C9.33332 17.84 10.08 19.5067 11.2853 20.7147L9.39999 22.6L9.19199 22.3853C7.62666 20.716 6.66666 18.4693 6.66666 16C6.66666 10.8453 10.8453 6.66667 16 6.66667ZM24.5227 12.1907C25.0427 13.3533 25.3333 14.644 25.3333 16C25.3333 18.5773 24.288 20.9107 22.6 22.6L20.7147 20.7147C21.92 19.5067 22.6667 17.84 22.6667 16C22.6667 15.4027 22.588 14.824 22.44 14.2733L24.5227 12.1907ZM21.656 8.45733L23.5427 10.3427L18.5773 15.3107C18.636 15.5307 18.6667 15.7613 18.6667 16C18.6667 17.4733 17.4733 18.6667 16 18.6667C14.5267 18.6667 13.3333 17.4733 13.3333 16C13.3333 14.5267 14.5267 13.3333 16 13.3333C16.2387 13.3333 16.4693 13.364 16.6893 13.4227L21.6573 8.45733H21.656Z"
              fill="white"
            />

        </svg>
      ),
    },
    {
      title: "Import spreadsheet",
      path: "/import-from-spreadsheet",
      icon: (
        <svg
          width="32"
          height="32"
          viewBox="0 0 32 32"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
 
            <path
              d="M17.6 16L21.3333 21.3333H18.1333L16 18.2853L13.8667 21.3333H10.6667L14.4 16L10.6667 10.6667H13.8667L16 13.7147L18.1333 10.6667H20V5.33333H6.66667V26.6667H25.3333V10.6667H21.3333L17.6 16ZM4 3.98933C4 3.25867 4.596 2.66667 5.332 2.66667H21.3333L28 9.33333V27.9907C28.0012 28.1658 27.968 28.3394 27.9021 28.5016C27.8362 28.6639 27.739 28.8115 27.6161 28.9362C27.4931 29.0609 27.3468 29.1601 27.1855 29.2283C27.0242 29.2964 26.8511 29.3321 26.676 29.3333H5.324C4.97384 29.3309 4.63869 29.1908 4.39096 28.9433C4.14322 28.6958 4.00279 28.3608 4 28.0107V3.98933Z"
              fill="white"
            />
    
        </svg>
      ),
    },
  ];
  const menuItens = routes.map((r, index) => (
    <Link key={index} to={r.path}>
      <div style={{ opacity: location === r.path ? '1' : '.5' }}>
      {r.icon}

      </div>
    </Link>
  ));
  return (
    <div className="sideBar">
      <div>
        <svg
          width="32"
          height="32"
          viewBox="0 0 32 32"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <path
            d="M14.1521 28C12.8973 28 11.6312 27.9548 10.3536 27.8644C9.07605 27.7966 7.95817 27.6384 7 27.3898V4.57627C7.5019 4.48588 8.03802 4.40678 8.60837 4.33898C9.17871 4.24859 9.76046 4.18079 10.3536 4.13559C10.9468 4.0904 11.5285 4.0565 12.0989 4.0339C12.692 4.0113 13.251 4 13.7757 4C15.2129 4 16.5475 4.11299 17.7795 4.33898C19.0114 4.54237 20.0722 4.89266 20.962 5.38983C21.8745 5.88701 22.5817 6.54237 23.0837 7.35593C23.5856 8.16949 23.8365 9.17514 23.8365 10.3729C23.8365 11.4802 23.5627 12.4294 23.0152 13.2203C22.4905 14.0113 21.749 14.6441 20.7909 15.1186C22.2281 15.5932 23.289 16.2938 23.9734 17.2203C24.6578 18.1469 25 19.322 25 20.7458C25 23.1638 24.1103 24.9831 22.3308 26.2034C20.5513 27.4011 17.8251 28 14.1521 28ZM11.2433 17.1525V24.3051C11.7224 24.3503 12.2357 24.3842 12.7833 24.4068C13.3308 24.4294 13.8327 24.4407 14.289 24.4407C15.1787 24.4407 16 24.3842 16.7529 24.2712C17.5285 24.1582 18.1901 23.9661 18.7376 23.6949C19.308 23.4011 19.7529 23.017 20.0722 22.5424C20.4144 22.0678 20.5856 21.4576 20.5856 20.7119C20.5856 19.3785 20.0951 18.452 19.1141 17.9322C18.1331 17.4124 16.7757 17.1525 15.0418 17.1525H11.2433ZM11.2433 13.7966H14.289C15.9316 13.7966 17.2205 13.5706 18.1559 13.1186C19.0913 12.6441 19.5589 11.8079 19.5589 10.6102C19.5589 9.48023 19.0684 8.67797 18.0875 8.20339C17.1293 7.72881 15.8745 7.49153 14.3232 7.49153C13.6616 7.49153 13.0684 7.50283 12.5437 7.52542C12.0418 7.54802 11.6084 7.58192 11.2433 7.62712V13.7966Z"
            fill="white"
            fillOpacity="0.5"
          />
        </svg>
      </div>
      <div className="sideBar--menu">{menuItens}</div>
      <div>
        <svg
          width="32"
          height="32"
          viewBox="0 0 32 32"
          fill="none"
          xmlns="http://www.w3.org/2000/svg"
        >
          <g opacity="0.5">
            <path
              d="M16 29.3333C8.63599 29.3333 2.66666 23.364 2.66666 16C2.66666 8.63599 8.63599 2.66666 16 2.66666C23.364 2.66666 29.3333 8.63599 29.3333 16C29.3333 23.364 23.364 29.3333 16 29.3333ZM14.6667 20V22.6667H17.3333V20H14.6667ZM17.3333 17.8067C18.4049 17.4837 19.3248 16.7866 19.9255 15.8423C20.5263 14.898 20.7679 13.7694 20.6063 12.6619C20.4447 11.5545 19.8908 10.5419 19.0453 9.8086C18.1999 9.07529 17.1192 8.67004 16 8.66666C14.9211 8.66657 13.8755 9.04019 13.041 9.72398C12.2066 10.4078 11.6347 11.3595 11.4227 12.4173L14.0387 12.9413C14.1129 12.5699 14.2911 12.2272 14.5525 11.9531C14.8139 11.679 15.1477 11.4847 15.5152 11.3929C15.8827 11.3011 16.2687 11.3156 16.6283 11.4346C16.9879 11.5536 17.3063 11.7723 17.5464 12.0652C17.7866 12.3581 17.9387 12.7132 17.9849 13.0891C18.0311 13.4651 17.9697 13.8464 17.8077 14.1888C17.6457 14.5311 17.3898 14.8205 17.0698 15.0231C16.7497 15.2257 16.3788 15.3333 16 15.3333C15.6464 15.3333 15.3072 15.4738 15.0572 15.7238C14.8071 15.9739 14.6667 16.313 14.6667 16.6667V18.6667H17.3333V17.8067Z"
              fill="white"
            />
          </g>
        </svg>
      </div>
    </div>
  );
};

export default Sidebar;
