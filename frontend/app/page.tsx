import Link from "next/link"


export default function Home() {
  return (
    <div>
      <Link href={"/register"}>
          <h1>
          Register
        </h1>
      </Link>
    </div>
  );
}
